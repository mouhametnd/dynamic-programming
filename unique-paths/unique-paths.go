package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		// {1, 5, 3, 5},
		// {4, 3, 4, 2},
		// {7, 2, 9, 1},

		// {0, 2, 2, 1},
		// {3, 1, 1, 1},
		// {4, 4, 2, 0},

		// {1, 1, 1},
		// {1, 0, 1},
		// {1, 1, 1},

		{1, 3, 2},
		{2, 5, 6},
		{3, 4, 1},
	}

	// problem #1
	{ /*
			Problem:
			Unique Paths
			A robot is located at the top-left corner of a m x n grid (marked 'S' in the diagram below).
			The robot can only move either down or right at any point in time.
			The robot is trying to reach the bottom-right corner of the grid (marked 'E' in the diagram below).
			How many possible unique paths are there?
			+---+---+---+---+
			| S |   |   |   |
			+---+---+---+---+
			|   |   |   |   |
			+---+---+---+---+
			|   |   |   | E |
			+---+---+---+---+
			Above is a 3 x 4 grid. How many possible unique paths are there?
		*/

		// Time complexity: m*n
		// Space complexity: m*n
		// F(i,j) = F(i-1,j) + F(i,j-1)
		uniquePaths := func(m int, n int) int {
			dp := make([][]int, m)
			for i := range dp {
				dp[i] = make([]int, n)
			}
			/*
				[
					[1,1,1],
					[1,2,3],
					[1,3,6],
					]
			*/
			dp[0][0] = 1
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					if i > 0 {
						dp[i][j] = dp[i-1][j]
					}
					if j > 0 {
						dp[i][j] += dp[i][j-1]
					}
				}
			}
			return dp[m-1][n-1]
		}
		// fmt.Println(uniquePaths(len(matrix), len(matrix[0])))
		uniquePaths(len(matrix), len(matrix[0]))
	}

	// problem #2
	{ /*
			Problem:
				Unique Paths with Obstales
				A robot is located at the top-left corner of a m x n grid (marked 'S' in the diagram below).
				The robot can only move either down or right at any point in time.
				The robot is trying to reach the bottom-right corner of the grid (marked 'E' in the diagram below).
				Now consider if some obstacles are added to the grids.
				How many unique paths would there be?
				+---+---+---+---+
				| S |   |   |   |
				+---+---+---+---+
				|   | 1 | 1 | 1 |
				+---+---+---+---+
				|   |   |   | E |
				+---+---+---+---+
				An obstacle and empty space is marked as 1 and 0 respectively in the grid.
		*/
		// Time complexity: O(m*n)
		// Space complexity: O(m*n)
		uniquePathsWithObstacles := func(grid [][]int) int {
			m := len(grid)
			n := len(grid[0])
			dp := make([][]int, m)
			for i := range dp {
				dp[i] = make([]int, n)
			}
			/*
				[
					[1,1,1]
					[1,0,1]
					[1,1,2]
				]
			*/
			dp[0][0] = 1
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					if grid[i][j] == 0 {
						continue
					}
					if i > 0 {
						dp[i][j] = dp[i-1][j]
					}
					if j > 0 {
						dp[i][j] += dp[i][j-1]
					}

				}
			}

			// fmt.Print(dp)
			return dp[m-1][n-1]
		}
		// fmt.Println(uniquePathsWithObstacles(matrix))
		uniquePathsWithObstacles(matrix)
	}

	// problem #3
	{
		/*
			Problem:
				Maximum Profit in a Grid
				A robot is located at the top-left corner of a m x n grid (marked 'S' in the diagram below).
				The robot can only move either down or right at any point in time.
				The robot is trying to reach the bottom-right corner of the grid (marked 'E' in the diagram below).
				Each cell contains a coin the robot can collect.
				What is the maximum profit the robot can accumulate?
				+---+---+---+---+
				| S | 2 | 2 | 1 |
				+---+---+---+---+
				| 3 | 1 | 1 | 1 |
				+---+---+---+---+
				| 4 | 4 | 2 | E |
				+---+---+---+---+
		*/

		// Time complexity: O(m*n)
		// Space complexity: O(m*n)
		// f(i,j) = max(f(i-1, j), f(i, j-1)) + grid(i,j)
		maxProfit := func(grid [][]int) int {
			m := len(grid)
			n := len(grid[0])
			dp := make([][]int, m)
			for i := range dp {
				dp[i] = make([]int, n)
			}
			/*
				[1,3,2]
				[2,5,6]
				[3,4,1]

				[1,4,6]
				[3,9,15]
				[6,13,16]
				y:2,1,1,0
				x:2,2,1,1,

				11,12,14,16,13
			*/
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {

					if i > 0 {
						dp[i][j] = dp[i-1][j]
					}
					if j > 0 && dp[i][j] < dp[i][j-1] {
						dp[i][j] = dp[i][j-1]
					}
					dp[i][j] += grid[i][j]
				}
			}
			cn := [][]int{}
			x := n - 1
			y := m - 1

			for x != -1 || y != -1 {
				cn = append(cn, []int{y, x})
				if x > 0 && y > 0 {
					if dp[y-1][x] > dp[y][x-1] {
						y--
					} else {
						x--
					}
					continue
				}

				if x > 0 {
					x--
				} else if y > 0 {
					y--
				} else {
					y--
					x--
				}
			}

			fmt.Println("path is", cn)
			return dp[m-1][n-1]
		}
		// fmt.Println(maxProfit(matrix))
		maxProfit(matrix)
	}

}
