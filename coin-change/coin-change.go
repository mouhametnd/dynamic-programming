package main

import (
	"fmt"
)

func main() {
	getCoins := func(n int) []int {
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
	// part 1: normal coin change
	{
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

		coinChange := func(n int) int {
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

		// fmt.Println(coinChange(3))
		// fmt.Println(coinChange(4))
		// fmt.Println(coinChange(6))
		// fmt.Println(coinChange(40))
		coinChange(40)
	}

	// part 2: coin change with exactly (t) coins
	{
		/*
			 Coin change with exactly (t) coins
			 Given an unlimited supply of coins of given denominations(j), find the total number of ways(n) to make a change of size n bu using exactly 9(?) coins.

			 denominations: 1,2,3,5
			 1. Objective function
				 f(n,t) return the different number of ways to make change of size n by using exactly t(9) coins

			 2. Base cases
				 f(i,0) = 0
				 f(0,0) = 1
				 f(i,1) = 1 if we have the denomination.

			 3. Transition function
				 f(n) += f(n-coin, t-1)

			 4. Excution order
				 Bottom-up

			 5. Answer location f(n,t)

			 time complexity: O(n)
			 space complexity: O(n)
		*/
		coinChange := func(n, t int) int {
			coins := []int{1, 2, 3, 5}
			dp := make([][]int, n+1)
			for i := range dp {
				dp[i] = make([]int, t+1)
			}
			dp[0][0] = 1
			for i := 0; i <= n; i++ {
				for j := 0; j <= t; j++ {
					if i != 0 && j != 0 {
						for _, coin := range coins {
							if i-coin >= 0 {
								dp[i][j] += dp[i-coin][j-1]
							}
						}
					}
				}
			}
			return dp[n][t]
		}

		// fmt.Println(coinChange(0, 0))
		// fmt.Println(coinChange(7, 3))
		coinChange(0, 0)
	}

	// part 2: coin change with no more than (t) coins
	{
		/*
			 Coin change with no more than (t) coins:
			 Given an unlimited supply of coins of given denominations, find the total number of ways to make a change of size n with no more than t coins.

			 denominations: 1,3,5,10

			 1. Objective function
			 f(n,t) return the num of ways to make a change of n by using no more than t coins.

			 2. Base cases
			 f(i,0) = 0
			 f(i,1) = 1 if we have the denomination.
			 f(0,1nyNum) = 1, because 0 is acceptable as t.

			 3. Transition function
			 f(n) += f(i-coin,j)

			 4. Execution order
			 bottom up

			 5. Answer location
				 f(n,t)
		*/

		coinChange := func(n, t int) int {
			coins := []int{1, 3, 5, 10}
			dp := make([][]int, n+1)
			for i := range dp {
				dp[i] = make([]int, t+1)
			}

			for i := 0; i <= n; i++ {
				for j := 0; j <= t; j++ {
					if i == 0 {
						dp[i][j] = 1
						continue
					}
					for _, coin := range coins {
						if i-coin >= 0 {
							dp[i][j] += dp[i-coin][j]
						}
					}
				}

			}

			return dp[n][t]
		}

		// fmt.Println(coinChange(3, 2))
		// fmt.Println(coinChange(10, 3))
		// fmt.Println()
		coinChange(0, 3)
	}

	// part 3:
	{
		/*
			 coin change:
				 Given an unlimited supply of coins of given denominations,
				 find the total number of ways to make a change of size n, by using an even number of coins.

				 denominations: 1,3,5,10

				 1. Objective function: Given n return the number of ways to make a change of size n with the given denomination by using even number of coins.
				 f(i,x) 
				 f(i,x) -> {
					x: 1 == even
					x: 0 == odd
				 }

				 2. Base Cases
				 f(0,0) = 0
				 f(0,1) = 1
				 f(1,0) = 1
				 f(1,1) = 0

				 f(2,0) = 0
				 f(2,1) = 1 -> (1*2)
				 f(3,0) = 2 -> (1*3), 3(1)
				 f(3,1) = 0
				 f(4,0) = 0
				 f(4,1) = 2 -> 1(1*4), 1(3 + 1),

				 3. Transition function
				 f(i,0) = f(i-coin, 1)
				 f(i,1) = f(i-coin, 0)

				 4. Execution order
					 Bottom up

				 5. Answer location
					 f(n)

		*/

		coinChange := func(n int) int {
			dp := make([][]int, n+1)
			for i := range dp {
				dp[i] = make([]int, 2)
			}
			dp[0][0] = 0
			dp[0][1] = 1
			// dp[1][0] = 1
			// dp[1][1] = 0

			// dp[2][0] = 0
			// dp[2][1] = 1

			// dp[3][0] = 2
			// dp[3][1] = 0

			// dp[4][0] = 0
			// dp[4][1] = 2

			// coins := []int{1,2,3,10}
			coins := getCoins(n)
			for i := 1; i <= n; i++ {
				for _, coin := range coins {
					if i-coin >= 0 {
						dp[i][0] += dp[i-coin][1]
						dp[i][1] += dp[i-coin][0]
					}
				}
			}

			return dp[n][1]
		}

		// fmt.Println(coinChange(0))
		// fmt.Println(coinChange(1))
		// fmt.Println(coinChange(2))
		// fmt.Println(coinChange(3))
		// fmt.Println(coinChange(4))
		// fmt.Println(coinChange(6))
		// fmt.Println(coinChange(4))
		fmt.Println(coinChange(7))
		fmt.Println(coinChange(8))

	}
}
