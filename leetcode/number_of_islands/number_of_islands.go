package number_of_islands

import "fmt"

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
