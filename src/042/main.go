package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://kyo-pro.hatenablog.com/entry/ABC172D
func solve(N int64) {
	sum := int64(0)
	for i := int64(1); i <= N; i++ {
		sum += (((N/i)*(N/i) + (N / i)) / 2) * i
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
	solve(N)
}
