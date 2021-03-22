package main

import "fmt"

func min(no1, no2 int) int {
	if no1 < no2 {
		return no1
	}

	return no2
}

// 最小路径 // 说明：每次只能向下或者向右移动一步。http://interview.wzcu.com/Leetcode/code/minimum_path_sum/
func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return 0
	}

	row := len(grid)
	col := len(grid[0])

	dp := make([][]int, row)
	for k := 0; k < row; k++ {
		dp[k] = make([]int, col)
	}

	dp[0][0] = grid[0][0] //caution

	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	for i := range dp {
		fmt.Println(dp[i])
	}

	return dp[row-1][col-1]
}

// 不同路径 // 说明：每次只能向下或者向右移动一步。http://interview.wzcu.com/Leetcode/code/unique_paths/
func uniquePathSum(m, n int) int {
	if m == 0 || n == 0 {
		return 1
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for j := 0; j < m; j++ {
		dp[j][0] = 1
	}

	for k := 0; k < n; k++ {
		dp[0][k] = 1
	}

	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			dp[x][y] = dp[x-1][y] + dp[x][y-1]
		}
	}

	return dp[m-1][n-1]
}

func main() {
	grid := [][]int{
		[]int{1,3,1},
		[]int{1,5,1},
		[]int{4,2,1},
	}

	fmt.Println(minPathSum(grid))

	fmt.Println(uniquePathSum(7, 3))
}
