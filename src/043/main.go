package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, M int64, A []int64, B []int64) {
	adj := adjacencyList(N, M, A, B)
	visited := make([]bool, N)
	dfs(0, adj, visited)
	for _, v := range visited {
		if !v {
			fmt.Println("The graph is not connected.")
			return
		}
	}
	fmt.Println("The graph is connected.")
}

func dfs(v int64, adj [][]int64, visited []bool) {
	visited[v] = true
	for _, u := range adj[v] {
		if !visited[u] {
			dfs(u, adj, visited)
		}
	}
}

func adjacencyList(N int64, M int64, A []int64, B []int64) (adj [][]int64) {
	adj = make([][]int64, N)
	for i := int64(0); i < M; i++ {
		adj[A[i]-1] = append(adj[A[i]-1], B[i]-1)
		adj[B[i]-1] = append(adj[B[i]-1], A[i]-1)
	}
	return
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
	var M int64
	scanner.Scan()
	M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, M)
	B := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, M, A, B)
}
