package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64) {
	dp := []int64{}
	for i := int64(0); i <= N+1; i++ {
		if i == 0 {
			dp = append(dp, 0)
		} else if i == 1 {
			dp = append(dp, 1)
		} else {
			dp = append(dp, dp[i-1]+dp[i-2])
		}
	}
	fmt.Println(dp[N+1])
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
