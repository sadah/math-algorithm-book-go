package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(N int64, h []int64) {
	dp := []int64{}
	for i := int64(0); i < N; i++ {
		if i == 0 {
			dp = append(dp, 0)
		} else if i == 1 {
			dp = append(dp, int64(math.Abs(float64(h[i]-h[i-1]))))
		} else {
			dp = append(dp,
				min(
					dp[i-1]+int64(math.Abs(float64(h[i]-h[i-1]))),
					dp[i-2]+int64(math.Abs(float64((h[i]-h[i-2])))),
				),
			)
		}
	}
	fmt.Println(dp[N-1])
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
	h := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		h[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, h)
}
