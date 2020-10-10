package dp

import (
	"fmt"
	"testing"
)

func Test_MultiMatrix(t *testing.T) {
	a := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}

	b := [][]int{
		[]int{1, 2},
		[]int{3, 4},
		[]int{5, 6},
	}

	res := [][]int{
		[]int{22, 28},
		[]int{49, 64},
	}


	fmt.Println(res)
	fmt.Println(multiMatrix(a, b))
}
