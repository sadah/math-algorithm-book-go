package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(N int64, K int64) {
	complementaryEvent := int64(0)
	for a := int64(1); a <= N; a++ {
		for b := max(1, a-(K-1)); b <= min(N, a+(K-1)); b++ {
			for c := max(1, a-(K-1)); c <= min(N, a+(K-1)); c++ {
				if int64(math.Abs(float64(b-c))) <= K-1 {
					complementaryEvent++
				}
			}
		}
	}
	fmt.Println(N*N*N - complementaryEvent)
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
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
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N, K)
}
