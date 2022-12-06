package main

import (
	"fmt"
)

/*
n = 3
k = 2
[0,3,2,4]
       21    '
[0,3,2,4,2.5,7,1]
      18,13

23, 21.

1. Objective function
	f(i) Return the minimum cost to reach the i-th index of the array with the possibility to jump 1 to k.
2. Base Cases
	f(0) = 0
	f(1) = 3
	f(2) = 2
	f(3) = 6 (2+4)
3. Recurrence relation
	f(n) = P(n) + min(f(n-1), f(n-2))
4. Excution order
	bottom-up
5. Answer location
	f(n)
*/

/*
Problem:
	Paid Staircase
	You are climbing a paid staircase. It takes n steps to reach to the top and you have to pay p[i] to step on the i-th stair. Each time you can climb 1 or 2 steps.
	What's the cheapest amount you have to pay to get to the top of the staircase?
*/

// Time complexity: O(n)
// Space complexity: O(n)
func paidStaircase(n int, p []int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = p[1]
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + p[i]
	}
	fmt.Println(dp)
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverser(arr []int) []int {
	leftPointer := 0
	rightPointer := len(arr) - 1
	for leftPointer < rightPointer {
		arr[leftPointer], arr[rightPointer] = arr[rightPointer], arr[leftPointer]
		leftPointer++
		rightPointer--
	}
	return arr
}
func getPathOfCheapeastPaidStaircase(n int, p []int) {
	dp := make([]int, n+1)
	from := make([]int, n+1)
	dp[0] = 0
	dp[1] = p[1]

	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + p[i]
		if dp[i-1] < dp[i-2] {
			from[i] = i - 1
		} else {
			from[i] = i - 2
		}
	}

	path := []int{0}
	for curr := n; curr > 0; curr = from[curr] {
		path = append(path, curr)
	}

	fmt.Println(path, "pathBefore")
	reverser(path)
	fmt.Println(path, "path")
	fmt.Println(dp[n], "cost")
}

func main() {
	stairs := []int{0, 3, 23, 4, 76, 5}
	println(paidStaircase(5, stairs))
	getPathOfCheapeastPaidStaircase(5, stairs)
}
