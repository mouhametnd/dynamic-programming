package main

import "fmt"
func pentFence(n int) int {
	/*
		   Problem:
		   	Paint Fence With Two Colors
		   	There is a fence with n posts, each post can be painted with either green or blue color.
		   	You have to paint all the posts such that no more than two adjacent fence posts have the same color.
		   	Return the total nber of ways you can paint the fence.

		## transition function
			f(i,j) = f(i-1,1-j) + f(i-2,1-j)

		## answer location
			f(n, 0) + f(n,1)
	*/

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[1][0] = 1 // +
	dp[1][1] = 1 // -
	dp[2][0] = 2 // +
	dp[2][1] = 2 // -              
	// dp[3][0] = 3 // -
	// dp[3][1] = 3 // +

	for i := 3; i <= n; i++ {
		for j := 0; j <= 1; j++ {
			dp[i][j] = dp[i-1][1-j] + dp[i-2][1-j]
		}
	}
	fmt.Println(dp)
	return dp[n][0] + dp[n][1]
}

func main() {
	fmt.Println( pentFence(99))
}
