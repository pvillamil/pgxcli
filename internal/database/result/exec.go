package result

import (
	"time"

	"github.com/jackc/pgx/v5/pgconn"
)

// ExecResult represents a non-row SQL command result (INSERT, UPDATE, DELETE, DDL).
type ExecResult struct {
	RowsAffected int64
	Status       string
	Duration     time.Duration
}

func NewExec(tag pgconn.CommandTag, duration time.Duration) *ExecResult {
	return &ExecResult{
		RowsAffected: tag.RowsAffected(),
		Status:       tag.String(),
		Duration:     duration,
	}
}

func (e *ExecResult) Type() Type {
	return ResultTypeExec
}
