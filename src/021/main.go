package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(n int64, r int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var n int64
    scanner.Scan()
    n, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var r int64
    scanner.Scan()
    r, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(n, r)
}
