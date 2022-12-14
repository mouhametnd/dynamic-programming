package main

/*

Problem:
	Climbing Stairs

	You are climbing a stair case. It takes n steps to reach to the top.
	Each time you can either climb 1 or 2 steps.
	In how many distinct ways can you climb to the top?

Framework for Solving DP Problems:
	1. Define the objective function
		f(n) is the number different ways to reach the i-th stair.
	2. Identify base cases
		f(0) = 1
		f(1) = 1
	3. Write down a recurrence relation for the optimized objective function
		f(n) = f(n-1) + f(n-2)
	4. What's the order of execution?
		bottom-up
	5. Where to look for the answer?
		f(n)


*/

// time complexity: O(n)
// space complexity: O(1)
func clmbStairs(n int) int {
	a := 1
	b := 1
	c := 2
	temp := 0
	for i := 3; i <= n; i++ {
		temp = a
		a = a + b + c
		b = temp
	}
	return a
}

// time complexity: O(n)
// space complexity: O(n)
func clmbStairs2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}

// time complexity: O(n*m)
// time complexity: O(n)
func clmbStairs3(n, k int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i-j != 0 {
				dp[i] += dp[j]
			}
		}

	}
	return dp[n]
}

func main() {
	println(clmbStairs(4))
	println(clmbStairs(5))
	println(clmbStairs2(4))
	println(clmbStairs3(4, 3))
	println(clmbStairs3(4, 2))
}
