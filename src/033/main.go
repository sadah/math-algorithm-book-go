package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(a_x int64, a_y int64, b_x int64, b_y int64, c_x int64, c_y int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var a_x int64
    scanner.Scan()
    a_x, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var a_y int64
    scanner.Scan()
    a_y, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var b_x int64
    scanner.Scan()
    b_x, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var b_y int64
    scanner.Scan()
    b_y, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var c_x int64
    scanner.Scan()
    c_x, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var c_y int64
    scanner.Scan()
    c_y, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(a_x, a_y, b_x, b_y, c_x, c_y)
}
