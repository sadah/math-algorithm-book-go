package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, K int64, x []int64, y []int64) {
	ans := int64(1 << 62)
	for i := int64(0); i < N; i++ {
		for j := int64(0); j < N; j++ {
			for k := int64(0); k < N; k++ {
				for l := int64(0); l < N; l++ {
					if checkPoint(N, x, y, x[i], x[j], y[k], y[l]) >= K {
						area := (x[i] - x[j]) * (y[k] - y[l])
						ans = min(ans, area)
					}
				}
			}
		}
	}
	fmt.Println(ans)
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func checkPoint(N int64, x []int64, y []int64, lx int64, rx int64, ly int64, ry int64) int64 {
	cnt := int64(0)
	for i := int64(0); i < N; i++ {
		if lx <= x[i] && x[i] <= rx && ly <= y[i] && y[i] <= ry {
			cnt++
		}
	}
	return cnt
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
	x := make([]int64, N)
	y := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, K, x, y)
}
