package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int) {
	k10, k5, k1 := 0, 0, 0
	k10 = N / 10000
	k5 = (N - 10000*k10) / 5000
	k1 = (N - 10000*k10 - 5000*k5) / 1000
	fmt.Println(k10 + k5 + k1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	solve(N)
}
