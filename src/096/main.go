package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(N int64, T []int64) {
	sumT := int64(0)
	for _, t := range T {
		sumT += t
	}

	dp := make([][]bool, N+1)
	for i := int64(0); i <= N; i++ {
		dp[i] = make([]bool, sumT+1)
	}
	dp[0][0] = true

	for i := int64(1); i <= N; i++ {
		for j := int64(0); j <= sumT; j++ {
			if j < T[i-1] {
				if dp[i-1][j] {
					dp[i][j] = true
					continue
				}
			}
			if j >= T[i-1] {
				if dp[i-1][j] || dp[i-1][j-T[i-1]] {
					dp[i][j] = true
				}
			}
		}
	}
	ans := int64(math.MaxInt64)
	for i := int64(0); i <= sumT; i++ {
		if dp[N][i] {
			cookTime := max(i, sumT-i)
			ans = min(ans, cookTime)
		}
	}
	fmt.Println(ans)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
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
	T := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		T[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, T)
}
