package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, W int64, w []int64, v []int64) {
	// Create a 2D array dp to store computed values for dynamic programming.
	dp := make([][]int64, N+1)
	for i := int64(0); i <= N; i++ {
		dp[i] = make([]int64, W+1)
	}

	// Loop through items (from 1 to N) and available weights (from 0 to W).
	for i := int64(1); i <= N; i++ {
		for j := int64(0); j <= W; j++ {
			// Initialize dp[i][j] with the value from the previous row.
			dp[i][j] = dp[i-1][j]

			// Check if the current item can fit into the available weight.
			if j >= w[i-1] {
				// Update dp[i][j] with the maximum value between:
				// 1. Not including the current item (dp[i-1][j]).
				// 2. Including the current item (dp[i-1][j-w[i-1]] + v[i-1]).
				dp[i][j] = max(dp[i][j], dp[i-1][j-w[i-1]]+v[i-1])
			}
		}
	}

	// The final answer is stored in dp[N][W], which represents the maximum value achievable
	// by selecting items from the list within the weight constraint W.
	fmt.Println(dp[N][W])
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var W int64
	scanner.Scan()
	W, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	w := make([]int64, N)
	v := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		w[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		v[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, W, w, v)
}
