package result

import "github.com/balaji01-4d/pgxspecial"

type SpecialResult struct {
	pgxspecial.SpecialCommandResult
}

func NewSpecial(result pgxspecial.SpecialCommandResult) *SpecialResult {
	return &SpecialResult{SpecialCommandResult: result}
}

func (s *SpecialResult) Type() ResultType {
	return ResultTypeSpecial
}
