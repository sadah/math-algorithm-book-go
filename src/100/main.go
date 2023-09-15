package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(Q int64, X []float64, Y []float64, Z []float64, T []int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var Q int64
    scanner.Scan()
    Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    X := make([]float64, Q)
    Y := make([]float64, Q)
    Z := make([]float64, Q)
    T := make([]int64, Q)
    for i := int64(0); i < Q; i++ {
        scanner.Scan()
        X[i], _ = strconv.ParseFloat(scanner.Text(), 64)
        scanner.Scan()
        Y[i], _ = strconv.ParseFloat(scanner.Text(), 64)
        scanner.Scan()
        Z[i], _ = strconv.ParseFloat(scanner.Text(), 64)
        scanner.Scan()
        T[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(Q, X, Y, Z, T)
}
