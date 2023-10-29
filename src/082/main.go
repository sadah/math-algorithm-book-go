package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(N int64, L []int64, R []int64) {
	time, ans := int64(0), int64(0)
	for {
		minEnd := int64(math.MaxInt64)
		for i := int64(0); i < N; i++ {
			if L[i] < time {
				continue
			}
			minEnd = min(minEnd, R[i])
		}
		if minEnd == math.MaxInt64 {
			break
		}
		time = minEnd
		ans++
	}
	fmt.Println(ans)
}

func min(x, y int64) int64 {
	if x < y {
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
	L := make([]int64, N)
	R := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		L[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		R[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, L, R)
}
