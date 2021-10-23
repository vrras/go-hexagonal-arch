package arithmetic

// Arith implements the Arithmetic interface
type Arith struct {
}

// New creates a new Arith
func New() *Arith {
	return &Arith{}
}

// Addition gets the result of adding parameters a and b
func (arith Arith) Addition(a, b int32) (int32, error) {
	return a + b, nil
}

// Substraction gets the result of substracting parameters a and b
func (arith Arith) Substraction(a, b int32) (int32, error) {
	return a - b, nil
}

// Multiplication gets the result of multiplying parameters a and b
func (arith Arith) Multiplication(a, b int32) (int32, error) {
	return a * b, nil
}

// Division gets the result of dividing parameters a and b
func (arith Arith) Division(a, b int32) (int32, error) {
	return a / b, nil
}
