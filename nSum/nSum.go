package main

/*
	Problem:
	Find the sum of N numbers;

	Objective function;
	f(1) is the sum of the first i elements.

	Recurrence relations;
	f(n) = f(n-1) + n;
*/

func nSumTopDown(n int) int {
	if n <= 1 {
		return n
	}
	return nSumTopDown(n-1) + n
}

func nSumBottomUp(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + i
	}
	return dp[n]
}

func main() {
	println("nSumTopDown", nSumTopDown(5))
	println("nSumBottomUp", nSumBottomUp(5))
}
