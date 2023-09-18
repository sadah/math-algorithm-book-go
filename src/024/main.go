package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, PQ [][]int64) {
	sum := float64(0)
	for i := int64(0); i < N; i++ {
		sum += float64(PQ[i][1]) / float64(PQ[i][0])
	}
	fmt.Println(sum)
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

	PQ := make([][]int64, N)

	for i := int64(0); i < N; i++ {
		PQ[i] = []int64{}
		scanner.Scan()
		p, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		PQ[i] = append(PQ[i], p)

		scanner.Scan()
		q, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		PQ[i] = append(PQ[i], q)
	}
	solve(N, PQ)
}
