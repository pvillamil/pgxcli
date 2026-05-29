package app

import (
	"context"
	"fmt"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/balajz/pgxcli/internal/app/renderer"
	"github.com/balajz/pgxcli/internal/app/ui"
	"github.com/balajz/pgxcli/internal/config"
	"github.com/balajz/pgxcli/internal/database"
	"github.com/balajz/pgxcli/internal/parser"
)

func (p *pgxCLI) runQuery(ctx context.Context, query string) tea.Msg {
	p.logger.Debug("executing query")
	stmts := parser.SplitSQLStatements(query)
	cmds := make([]tea.Cmd, 0, len(stmts))

StatementsLoop:
	for _, stmt := range stmts {
		p.logger.Debug("parsed statement", "statement", stmt)
		if stmt == "" || stmt == ";" {
			continue
		}

		start := time.Now()
		queryResult, _, err := p.client.ExecuteQuery(ctx, stmt, false)
		execDuration := time.Since(start)
		if err != nil {
			p.logger.Error("query execution failed", "error", err)
			cmds = append(cmds, p.printError(err))
			if p.config.Main.OnError == config.OnErrorStop {
				break StatementsLoop
			}
			continue
		}
		resultCmd, err := p.handleQueryResult(queryResult, execDuration)
		if err != nil {
			p.logger.Error("error handling query result", "error", err)
			cmds = append(cmds, p.printError(err))
			if p.config.Main.OnError == config.OnErrorStop {
				break StatementsLoop
			}
			continue
		}
		cmds = append(cmds, resultCmd)
	}

	return ui.ExecCmdMsg{Cmd: p.withPrompt(cmds...)}
}

func (p *pgxCLI) handleQueryResult(r database.Rows, execDuration time.Duration) (cmd tea.Cmd, err error) {
	var s strings.Builder

	cols := renderer.GetColumnStrings(r, true)
	if len(cols) > 0 {
		rowIter := renderer.NewRowIter(r, true)
		if err := renderer.TableRender(cols, rowIter, "", &s, &s, p.config); err != nil {
			r.Close() // Ensure closed on error
			return nil, err
		}
	}

	// We must close the rows before reading the tag
	if closeErr := r.Close(); closeErr != nil {
		return nil, closeErr
	}

	tag, err := r.Tag()
	if err != nil {
		return nil, err
	}
	tagStr := tag.String()
	if tagStr == "" {
		tagStr = "OK"
	}

	output := s.String()
	if len(cols) == 0 {
		output = tagStr
	} else {
		output += tagStr
	}

	// Append timing info to the output
	timingInfo := fmt.Sprintf("\nTime %.3fs", execDuration.Seconds())
	output += timingInfo

	return p.printViaPager(output), nil
}
