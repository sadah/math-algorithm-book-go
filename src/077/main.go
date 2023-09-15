package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, X []int64, Y []int64) {

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
