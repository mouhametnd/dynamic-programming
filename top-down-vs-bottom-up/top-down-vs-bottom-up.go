package main

import (
	"fmt"
	// "log"
	"time"
)

// this is NOT dynamic programming, but a regular recursion
// time complexity: O(2^n)
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// this is a top-down dynamic programming (a.k.a. recursion + memoization)
// time complexity: O(n)
func fibTopDown(n int) int {
	memo := map[int]int{}
	return fibTopDownHelperf(n, memo)
}

func fibTopDownHelperf(n int, memo map[int]int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = fibTopDownHelperf(n-1, memo) + fibTopDownHelperf(n-2, memo)
	return memo[n]
}

// this is a bottom-up dynamic programming (forward dynamic programming)
//
// f(i-1)
//
//	\
//	 >-------> f(i)
//	/
//
// f(i-2)
//
// time complexity: O(n)
func fibBottomUpDPForward(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// this is a bottom-up dynamic programming (backward dynamic programming)
//
//   -----> f(i+2)
//   |
// f(i)
//   |
//   -----> f(i+1)
// time complexity: O(n)
/*
	using the solution of a non-base case problem, we solve a new non-base sub-problem.
*/
func fibBottomUpDPBackward(n int) int {
	dp := make([]int, n+2)
	dp[0] = 0
	dp[1] = 1
	for i := 1; i < n; i++ {
		dp[i+1] += dp[i]
		dp[i+2] += dp[i]
	}
	return dp[n]
}

func main() {
	start := time.Now()
	// fibBottomUpDPForward(1000000) // 0.004s
	// fibBottomUpDPBackward(1000000) // 0.005s
	// fibTopDown(1000000) // .5s
	// fib(50) //39s
	// fibTopDown(10000)
	elapsed := time.Since(start).Seconds()
	fmt.Println("Binomial took", elapsed)
}
 
