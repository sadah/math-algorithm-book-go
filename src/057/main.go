package main

import (
	"bufio"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(K int64, N int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var K int64
    scanner.Scan()
    K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var N int64
    scanner.Scan()
    N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(K, N)
}
