package main

import (
	"fmt"
	"math"
)


/*

Problem:
	Shortest Path in DAG
	Given a graph as an adjacency matrix, find the path from the first to the last vertex(node).

	Objective function:
	F[i] is the shortest path from i to the last vertex.

	Transition function:
	F[i] = min[weight + F[j]]

	Base case:
	F[n] = 0
*/

func findShortestPath(graph [][]int) int {

	// fmt.Println("test", 4)
	

	if len(graph) == 0 {
		return 0
	}

	n := len(graph)
	dp := make([]int, n)

	for i := range dp {
		dp[i] = math.MaxInt64
	}
	dp[n-1] = 0

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(graph[i]); j++ {
			weight := graph[i][j]
			if weight > 0 && dp[i] > weight+dp[j] {
				dp[i] = weight + dp[j]

			}
		}
	}

	return dp[0]
}

func main() {
	var graph [][]int = make([][]int, 10)
	for i := range graph {
		graph[i] = make([]int, 10)
	}

	graph[0][1] = 2
	graph[0][2] = 2
	graph[0][3] = 3

	graph[1][4] = 2
	graph[1][5] = 1
	graph[1][6] = 1

	graph[2][4] = 1
	graph[2][5] = 2
	graph[2][6] = 1

	graph[3][4] = 2
	graph[3][5] = 3
	graph[3][6] = 2

	graph[4][7] = 3
	graph[4][8] = 2

	graph[5][7] = 3
	graph[5][8] = 3

	graph[6][7] = 1
	graph[6][8] = 3

	graph[7][9] = 4
	graph[8][9] = 3

	fmt.Println(findShortestPath(graph))

	// arr := []int{1, 2, 3, 4, 5}
	// arr[2] = arr[len(arr)-1] -> [1, 2, 5, 4, 5]
	// arr[:len(arr)-1] -> [1, 2, 5, 4]
	
}


