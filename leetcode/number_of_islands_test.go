package leetcode

import (
	"fmt"
	"testing"
)

/*
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:
```
Input:
11110
11010
11000
00000

Output: 1
```

Example 2:
```
Input:
11000
11000
00100
00011

Output: 3
```
*/

// https://www.youtube.com/watch?v=o8S2bO3pmO4

func numIslands(grid [][]byte) int {
	num := 0
	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("grid[%d][%d] = %d: ", i, j, grid[i][j])
			if grid[i][j] == '1' {
				num += dfs(grid, i, j)
			}
			fmt.Println()
		}
	}
	return num
}

func dfs(grid [][]byte, i, j int) int {
	if i <= -1 || i >= len(grid) || j <= -1 || j >= len(grid[i]) {
		return 0
	}
	fmt.Printf("dfs(%d, %d) = %d, ", i, j, grid[i][j])
	if grid[i][j] == '0' {
		return 0
	}

	// Mark as visited
	grid[i][j] = '0'

	dfs(grid, i+1, j) // down
	dfs(grid, i-1, j) // up
	dfs(grid, i, j+1) // right
	dfs(grid, i, j-1) // left

	fmt.Printf("finish, ")
	return 1
}

func TestNumIslands(t *testing.T) {
	tests := map[string]struct {
		input [][]byte
		want  int
	}{
		"example1": {
			input: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		"example2": {
			input: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := numIslands(test.input); got != test.want {
				t.Errorf("numIslands: got=%v, want=%v", got, test.want)
			}
		})
	}
}
