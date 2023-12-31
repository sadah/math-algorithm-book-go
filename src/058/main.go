package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(N int64, X int64, Y int64) {
	sum := int64(math.Abs(float64(X)) + math.Abs(float64(Y)))
	if sum > N {
		fmt.Println(NO)
		return
	}
	if (X+Y)%2 != N%2 {
		fmt.Println(NO)
		return
	}
	fmt.Println(YES)
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
	var X int64
	scanner.Scan()
	X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var Y int64
	scanner.Scan()
	Y, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N, X, Y)
}
