package result

type ResultType int

const (
	ResultTypeQuery ResultType = iota
	ResultTypeExec
	ResultTypeSpecial
)

// Result marks values returned by SQL execution paths.
type Result interface {
	Type() ResultType
}
