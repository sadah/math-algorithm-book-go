package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(a_x int64, a_y int64, b_x int64, b_y int64, c_x int64, c_y int64) {
	ba_x := b_x - a_x
	ba_y := b_y - a_y
	bc_x := b_x - c_x
	bc_y := b_y - c_y
	ca_x := c_x - a_x
	ca_y := c_y - a_y
	cb_x := c_x - b_x
	cb_y := c_y - b_y

	pattern := 2
	if ba_x*bc_x+ba_y*bc_y < 0 {
		pattern = 1
	}
	if ca_x*cb_x+ca_y*cb_y < 0 {
		pattern = 3
	}

	ans := float64(0)
	if pattern == 1 {
		ans = math.Sqrt(float64(ba_x*ba_x + ba_y*ba_y))
	} else if pattern == 3 {
		ans = math.Sqrt(float64(ca_x*ca_x + ca_y*ca_y))
	} else if pattern == 2 {
		ans = math.Abs(float64(ba_x*ca_y-ba_y*ca_x)) / math.Sqrt(float64(bc_x*bc_x+bc_y*bc_y))
	}
	fmt.Println(ans)
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
