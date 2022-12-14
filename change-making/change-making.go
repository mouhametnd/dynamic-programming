package main

import (
	"fmt"
	"math"
)

/*
	Change-making Problem
	Given an unlimited supply of coins of given denominations,
	what is the minimum number of coins needed to make a change of size n?
	coins = 1, 3, 5

	If we don;t have any chances, let use the world logic .

	1. Objetive function
		f(n) return the minimum number of coins to make change of size n.

	2. Base cases
		f(0) = 0 (0)
		f(1) = 1 (1)

	4. Transition Function
	f(n) = min(n-currCoint) +1


*/

func changeMaking(n int) int {
	dp := make([]int, n+1)
	// coins := []int{1, 3, 5}
	coins := []int{1, 4, 5}
	dp[1] = 1
	min := math.MaxInt
	for i := 2; i <= n; i++ {
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin]< min{
				min = dp[i-coin]
			}
		}
		dp[i] = min + 1
		min = math.MaxInt
	}
	return dp[n]
}

func main() {
	// fmt.Println(changeMaking(3))
	fmt.Println(changeMaking(8))
}
