package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, K int64, V []int64) {

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
    V := make([]int64, K)
    for i := int64(0); i < K; i++ {
        scanner.Scan()
        V[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, K, V)
}
