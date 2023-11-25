package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, B []int64) {
	ans := B[0] + B[N-2]
	for i := int64(1); i < N-1; i++ {
		ans += min(B[i-1], B[i])
	}
	fmt.Println(ans)
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
	B := make([]int64, N-1)
	for i := int64(0); i < N-1; i++ {
		scanner.Scan()
		B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, B)
}
