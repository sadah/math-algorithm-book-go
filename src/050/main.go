package main

import (
	"bufio"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(a int64, b int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var a int64
    scanner.Scan()
    a, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var b int64
    scanner.Scan()
    b, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(a, b)
}
