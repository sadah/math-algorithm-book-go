package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, M int64, A []int64, B []int64) {
	g := make([][]int64, N+1)
	for i := int64(0); i < M; i++ {
		g[A[i]] = append(g[A[i]], B[i])
		g[B[i]] = append(g[B[i]], A[i])
	}

	dist := make([]int64, N+1)
	for i := range dist {
		dist[i] = -1
	}

	queue := make([]int64, 0, N)
	queue = append(queue, 1)
	dist[1] = 0

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		for _, v := range g[pos] {
			if dist[v] == -1 {
				dist[v] = dist[pos] + 1
				queue = append(queue, v)
			}
		}
	}

	for i := int64(1); i <= N; i++ {
		if 0 <= dist[i] && dist[i] <= 120 {
			fmt.Println(dist[i])
		} else {
			fmt.Println(120)
		}
	}
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
