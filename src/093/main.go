package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(A int64, B int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var A int64
    scanner.Scan()
    A, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var B int64
    scanner.Scan()
    B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(A, B)
}
