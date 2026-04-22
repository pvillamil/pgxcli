// Package parser provides SQL classification and statement splitting helpers.
package parser

import (
	"strings"

	pg_query "github.com/pganalyze/pg_query_go/v6"
)

const (
	commandTypeInvalid = "INVALID"
	commandTypeQuery   = "QUERY"
	commandTypeExecute = "EXECUTE"
)

// CommandType classifies SQL text as QUERY, EXECUTE, or INVALID.
func CommandType(sql string) string {
	trimmed := strings.TrimSpace(sql)
	if trimmed == "" {
		return commandTypeInvalid
	}

	if result, ok := classifyFast(trimmed); ok {
		return result
	}

	return classifyWithAST(trimmed)
}

// classifyFast handles the common single-statement cases without invoking the
// pg_query C parser. It returns ("", false) to signal "defer to AST" whenever
// the input is ambiguous:
//   - SQL starting with a line/block comment (firstKeyword returns "--" or "")
//   - SELECT containing INTO anywhere, including in string literals
//   - EXECUTE statements (prepared stmt body determines write-ness, not keyword)
//   - Multi-statement input (inner semicolon detected)
//
// These are conservative fallbacks, not bugs.
func classifyFast(sql string) (string, bool) {
	if hasInnerSemicolon(sql) {
		return "", false
	}

	keyword := firstKeyword(sql)
	switch keyword {
	case "SELECT":
		if hasToken(sql, "INTO") {
			return "", false
		}
		return commandTypeQuery, true
	case "SHOW", "EXPLAIN", "TABLE", "VALUES", commandTypeExecute:
		return commandTypeQuery, true
	case "BEGIN", "COMMIT", "ROLLBACK", "SAVEPOINT", "RELEASE", "START",
		"SET", "VACUUM", "ANALYZE", "REINDEX", "GRANT", "REVOKE",
		"CREATE", "ALTER", "DROP", "TRUNCATE", "MERGE":
		return commandTypeExecute, true
	default:
		return "", false
	}
}

func hasInnerSemicolon(sql string) bool {
	trimmed := strings.TrimSpace(sql)
	trimmed = strings.TrimRight(trimmed, ";")
	return strings.Contains(trimmed, ";")
}

func firstKeyword(sql string) string {
	fields := strings.Fields(sql)
	if len(fields) == 0 {
		return ""
	}
	keyword := strings.Trim(fields[0], "();")
	return strings.ToUpper(keyword)
}

func hasToken(sql string, token string) bool {
	parts := strings.FieldsFunc(sql, func(r rune) bool {
		switch r {
		case ' ', '\n', '\r', '\t', '\f', '\v', ',', ';', '(', ')':
			return true
		default:
			return false
		}
	})

	for _, part := range parts {
		if strings.EqualFold(part, token) {
			return true
		}
	}
	return false
}

func classifyWithAST(sql string) string {
    tree, err := pg_query.Parse(sql)
    if err != nil {
        return commandTypeInvalid
    }
    if len(tree.Stmts) == 0 {
        return commandTypeInvalid
    }

    for _, stmt := range tree.Stmts {
        if nodeWrites(stmt.Stmt.Node) {
            return commandTypeExecute
        }
    }
    return commandTypeQuery
}

//nolint:cyclop // dispatch table over a closed AST node set; cases are irreducible
func nodeWrites(node any) bool {
    switch n := node.(type) {
    case *pg_query.Node_SelectStmt:
        return n.SelectStmt.IntoClause != nil
    case *pg_query.Node_InsertStmt:
        return len(n.InsertStmt.ReturningList) == 0
    case *pg_query.Node_UpdateStmt:
        return len(n.UpdateStmt.ReturningList) == 0
    case *pg_query.Node_DeleteStmt:
        return len(n.DeleteStmt.ReturningList) == 0
    case *pg_query.Node_VariableShowStmt,
        *pg_query.Node_ExplainStmt,
        *pg_query.Node_ExecuteStmt:
        return false
    case *pg_query.Node_CreateStmt,
        *pg_query.Node_AlterTableStmt,
        *pg_query.Node_DropStmt,
        *pg_query.Node_TruncateStmt,
        *pg_query.Node_RenameStmt,
        *pg_query.Node_VariableSetStmt:
        return true
    case *pg_query.Node_CopyStmt:
        return n.CopyStmt.IsFrom
    default:
        return true
    }
}

// IsQuery returns true if the SQL statement is a read-only query.
func IsQuery(sql string) bool {
	return CommandType(sql) == commandTypeQuery
}

// IsExecute returns true if the SQL statement modifies data.
func IsExecute(sql string) bool {
	return CommandType(sql) == commandTypeExecute
}

// IsValid returns true if the SQL statement can be parsed successfully.
func IsValid(sql string) bool {
	return classifyWithAST(strings.TrimSpace(sql)) != commandTypeInvalid
}
