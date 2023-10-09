package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(A int64, B int64) {
	ans := int64(0)
	for i := int64(1); i <= B; i++ {
		if solver(A, B, i) {
			ans = i
		}
	}
	fmt.Println(ans)
}

func solver(A, B, t int64) bool {
	cl := (A + t - 1) / t
	cr := B / t
	return cr-cl >= 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var A int64
	scanner.Scan()
	A, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var B int64
	scanner.Scan()
	B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(A, B)
}
