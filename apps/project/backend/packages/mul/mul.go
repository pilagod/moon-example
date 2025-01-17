package mul

import "github.com/pilagod/moon-example/project/backend/packages/add"

func Mul(a, b int) int {
	result := 0
	for i := 0; i < b; i++ {
		result = add.Add(result, a)
	}
	return result
}
