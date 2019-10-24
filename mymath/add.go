package mymath

// Add adds the input integer.
// BUG(mepaling): can only used in "int" type
func Add(a int, b int) int {
	return a + b
}

// OldAdd adds the input integer.
//
// Deprecated: Use Add for instead.
func OldAdd(a, b int) int {
	return a + b
}
