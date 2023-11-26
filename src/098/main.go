package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, X []int64, Y []int64, A int64, B int64) {
	count := 0
	for i := int64(0); i < N; i++ {
		xa := X[i] - A
		ya := Y[i] - B
		xb := X[(i+1)%N] - A
		yb := Y[(i+1)%N] - B
		if ya > yb {
			tmpA := xa
			xa = xb
			xb = tmpA
			tmpB := ya
			ya = yb
			yb = tmpB
		}
		if ya <= 0 && 0 < yb && xa*yb-xb*ya < 0 {
			count += 1
		}
	}
	if count%2 == 1 {
		fmt.Println("INSIDE")
	} else {
		fmt.Println("OUTSIDE")
	}
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
	X := make([]int64, N)
	Y := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		X[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		Y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	var A int64
	scanner.Scan()
	A, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var B int64
	scanner.Scan()
	B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N, X, Y, A, B)
}
