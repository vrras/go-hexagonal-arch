package api

// Arithmetic
type Arithmetic interface {
	Addition(a, b int32) (int32, error)
	Substraction(a, b int32) (int32, error)
	Multiplication(a, b int32) (int32, error)
	Division(a, b int32) (int32, error)
}
