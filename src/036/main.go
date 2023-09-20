package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(A int64, B int64, H int64, M int64) {
	rad := math.Pi * 2 * (float64(H)/12 + float64(M)/60/12 - float64(M)/60)
	ans := math.Sqrt(float64(A*A+B*B) - 2*float64(A*B)*math.Cos(rad))
	fmt.Println(ans)
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
	var H int64
	scanner.Scan()
	H, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var M int64
	scanner.Scan()
	M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(A, B, H, M)
}
