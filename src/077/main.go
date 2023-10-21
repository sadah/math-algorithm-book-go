package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, X []int64, Y []int64) {
	ans := int64(0)

	sort.Slice(X, func(i, j int) bool {
		return X[i] < X[j]
	})

	sort.Slice(Y, func(i, j int) bool {
		return Y[i] < Y[j]
	})

	for i := int64(0); i < N; i++ {
		ans += X[i] * (-N + 2*(i+1) - 1)
		ans += Y[i] * (-N + 2*(i+1) - 1)
	}
	fmt.Println(ans)
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
	X := make([]int64, N)
	Y := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		X[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		Y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, X, Y)
}
