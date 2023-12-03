package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, c string) {
	ans := int64(0)
	for i := int64(0); i < N; i++ {
		code := int64(0)
		if c[i:i+1] == "B" {
			code = 0
		}
		if c[i:i+1] == "W" {
			code = 1
		}
		if c[i:i+1] == "R" {
			code = 2
		}
		ans += code * ncr(N-1, i)
		ans %= 3
	}
	if N%2 == 0 {
		ans = (3 - ans) % 3
	}
	fmt.Println("BWR"[ans : ans+1])
}

// リュカの定理で ncr mod 3 を計算
func ncr(x, y int64) int64 {
	if x < 3 && y < 3 {
		if x < y {
			return 0
		}
		if x == 2 && y == 1 {
			return 2
		}
		return 1
	}
	return ncr(x/3, y/3) * ncr(x%3, y%3) % 3
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
	scanner.Scan()
	c := scanner.Text()
	solve(N, c)
}
