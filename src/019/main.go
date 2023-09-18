package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, A []int64) {
	m := make(map[int64]int64)
	for k := range A {
		_, ok := m[A[k]]
		if ok {
			m[A[k]]++
		} else {
			m[A[k]] = 1
		}
	}
	sum := int64(0)
	sum += combination(m[1], 2) + combination(m[2], 2) + combination(m[3], 2)
	fmt.Println(sum)
}

func combination(n, r int64) int64 {
	if n < r {
		return 0
	}
	if n == r {
		return 1
	}
	if r == 1 {
		return n
	}
	return combination(n-1, r) + combination(n-1, r-1)
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
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A)
}
