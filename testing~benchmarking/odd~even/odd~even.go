package even

// Test for Even
func Even(i int) bool { // Exported functions
	return i%2 == 0
}

// Test for Odd
func Odd(i int) bool {
	return i%2 != 0
}
