package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, Q int64, L []int64, R []int64, X []int64) {
	differences := make([]int64, N+1)
	for i := int64(0); i < Q; i++ {
		differences[L[i]-1] += X[i]
		differences[R[i]] -= X[i]
	}
	for i := int64(1); i < N; i++ {
		if differences[i] == 0 {
			fmt.Print("=")
		} else if differences[i] > 0 {
			fmt.Print("<")
		} else {
			fmt.Print(">")
		}
	}
	fmt.Println()
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
	var Q int64
	scanner.Scan()
	Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	L := make([]int64, Q)
	R := make([]int64, Q)
	X := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		L[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		R[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		X[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, Q, L, R, X)
}
