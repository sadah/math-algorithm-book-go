package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, a []int64, b []int64) {
	g := make([][]int64, N+1)
	for i := int64(0); i < N-1; i++ {
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	dp := make([]int64, N+1)
	visited := make([]bool, N+1)
	dp = dfs(1, g, visited, dp)

	ans := int64(0)
	for i := int64(2); i <= N; i++ {
		ans += dp[i] * (N - dp[i])
	}

	fmt.Println(ans)
}

func dfs(pos int64, g [][]int64, visited []bool, dp []int64) []int64 {
	visited[pos] = true
	dp[pos] = 1
	for _, v := range g[pos] {
		if !visited[v] {
			dp = dfs(int64(v), g, visited, dp)
			dp[pos] += dp[v]
		}
	}
	return dp
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
	a := make([]int64, N-1)
	b := make([]int64, N-1)
	for i := int64(0); i < N-1; i++ {
		scanner.Scan()
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		b[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, a, b)
}
