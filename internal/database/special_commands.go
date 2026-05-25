package database

import (
	"context"

	"github.com/balajz/pgxcli/pgxspecial"
	"github.com/balajz/pgxcli/pgxspecial/database"
	// Register built-in pgxspecial commands via package init side effects.
	_ "github.com/balajz/pgxcli/pgxspecial/dbcommands"
)

func init() {
	registerSpecialCommands()
}

func registerSpecialCommands() {
	pgxspecial.RegisterCommand(pgxspecial.SpecialCommandRegistry{
		Cmd:         "\\q",
		Alias:       []string{"\\quit", "\\exit"},
		Syntax:      "\\q",
		Description: "Quit Pgxcli",
		Handler: func(_ context.Context, _ database.Queryer, _ string, _ bool) (pgxspecial.SpecialCommandResult, error) {
			return ExitAction{}, nil
		},
		CaseSensitive: false,
	})

	pgxspecial.RegisterCommand(pgxspecial.SpecialCommandRegistry{
		Cmd:         "\\c",
		Syntax:      "\\c database_name",
		Description: "Change a new database",
		Handler: func(_ context.Context, _ database.Queryer, s string, _ bool) (pgxspecial.SpecialCommandResult, error) {
			return ChangeDbAction{Name: s}, nil
		},
		CaseSensitive: true,
		Alias:         []string{"\\connect"},
	})

	pgxspecial.RegisterCommand(pgxspecial.SpecialCommandRegistry{
		Cmd:         "\\conninfo",
		Syntax:      "\\conninfo",
		Description: "Get connection details",
		Handler: func(_ context.Context, _ database.Queryer, _ string, _ bool) (pgxspecial.SpecialCommandResult, error) {
			return ConnInfoAction{}, nil
		},
		CaseSensitive: false,
	})
}

// ExitAction indicates that the REPL should terminate.
type ExitAction struct{}

// ChangeDbAction carries target database name for \c / \connect commands.
type ChangeDbAction struct {
	Name string
}

// ConnInfoAction indicates that connection info should be displayed.
type ConnInfoAction struct{}
