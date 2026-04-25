package result

type Type int

const (
	ResultTypeQuery Type = iota
	ResultTypeExec
	ResultTypeSpecial
)

// Result marks values returned by SQL execution paths.
type Result interface {
	Type() Type
}
