package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, W int64, w []int64, v []int64) {

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
    var W int64
    scanner.Scan()
    W, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    w := make([]int64, N)
    v := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        w[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        v[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, W, w, v)
}
