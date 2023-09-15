package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(H int64, W int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var H int64
    scanner.Scan()
    H, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var W int64
    scanner.Scan()
    W, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(H, W)
}
