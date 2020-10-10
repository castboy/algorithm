package dp

func multiMatrix(a, b [][]int) [][]int {
	line := len(a)
	col := len(b[0])
	res := make([][]int, line)

	for l := 0; l < line; l++ {
		res[l] = make([]int,col)
		for c := 0; c < col; c++ {
			res[l][c] += a[l][c] * b[c][l]
		}
	}

	return res
}

