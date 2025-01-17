package sub

func Sub(a, b int) int {
	if a < b {
		panic("a must be greater than b")
	}
	return a - b
}
