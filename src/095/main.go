package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, C []int64, P []int64, Q int64, L []int64, R []int64) {

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
    C := make([]int64, N)
    P := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        C[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        P[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    var Q int64
    scanner.Scan()
    Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    L := make([]int64, Q)
    R := make([]int64, Q)
    for i := int64(0); i < Q; i++ {
        scanner.Scan()
        L[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        R[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, C, P, Q, L, R)
}
