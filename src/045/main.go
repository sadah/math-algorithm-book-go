package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, M int64, a []int64, b []int64) {
	adj := adjacencyList(N, M, a, b)
	sum := int64(0)
	for i, v := range adj {
		count := 0
		for _, u := range v {
			if int64(i) > u {
				count++
			}
		}
		if count == 1 {
			sum++
		}
	}
	fmt.Println(sum)
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
	a := make([]int64, M)
	b := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		b[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, M, a, b)
}
