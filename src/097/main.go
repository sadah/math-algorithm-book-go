package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(L int64, R int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var L int64
    scanner.Scan()
    L, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var R int64
    scanner.Scan()
    R, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(L, R)
}
