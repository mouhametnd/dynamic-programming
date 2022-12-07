package main

import "fmt"

/*
Problem:
	Coin change

	Given an unlimited supply of coins of given denominations(j), find the total number of ways(n) to make a change of size n.

	denominations: 1,3,5,10
	1. Objective function
		f(i) return the different number of ways to make change of size n.

	2. Base cases
		f(0) = 1;

	3. transition function
		f(n) += f(n-j(coin))'

	4. Excution order
		Bottom-up

	5. answer location
		f(n)

	Implemented by knowing our base cases, func boundaries, transition function, execution order and answer location.

	time complexity: O(n)
	space complexity: O(n)
*/
func getCoins(n int) []int {
	coins := []int{1}
	if n >= 10 {
		return append(coins, 3, 5, 10)
	}
	if n >= 5 {
		return append(coins, 3, 5)
	}
	if n >= 3 {
		return append(coins, 3)
	}
	return coins
}
func coinChange(n int) int {
	coins := getCoins(n)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[n]
}

func coinChange2(n int) int {
	/*
		Coin change
					 				   
	Given an unlimited supply of coins of given denominations(j), find the total number of ways(n) to make a change of size n bu using exactly 9(?) coins.

	denominations: 1,3,5,10
	1. Objective function
		f(n) return the different number of ways to make change of size n by using exactly 9 coins
		
		2. Base cases
		input must be bigger than 8.

	4. Excution order
		Bottom-up

	Implemented by knowing our base cases, func boundaries, transition function, execution order and answer location.

	time complexity: O(n)
	space complexity: O(n)
	*/
	coins := getCoins(n)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[n]
}

func main() {
	// fmt.Println(coinChange(3))
	// fmt.Println(coinChange(4))
	// fmt.Println(coinChange(6))
	fmt.Println(coinChange(40))
}
