package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, a []int64, b []int64, c []int64) {

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
    a := make([]int64, N)
    b := make([]int64, N)
    c := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        b[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        c[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, a, b, c)
}
