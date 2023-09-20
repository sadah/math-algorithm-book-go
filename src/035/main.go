package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(x []int64, y []int64, r []int64) {
	x1, y1, r1 := x[0], y[0], r[0]
	x2, y2, r2 := x[1], y[1], r[1]

	// [1]　一方の円が他方の円を完全に含み、2つの円は接していない
	if (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) < (r1-r2)*(r1-r2) {
		fmt.Println(1)
		return
	}
	// [2]　一方の円が他方の円を完全に含み、2つの円は接している
	if (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) == (r1-r2)*(r1-r2) {
		fmt.Println(2)
		return
	}
	// [3]　2つの円が互いに交差する
	if (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) < (r1+r2)*(r1+r2) {
		fmt.Println(3)
		return
	}
	// [4]　2つの円の内部に共通部分は存在しないが、2つの円は接している
	if (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) == (r1+r2)*(r1+r2) {
		fmt.Println(4)
		return
	}
	// [5]　2つの円の内部に共通部分は存在せず、2つの円は接していない
	fmt.Println(5)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	x := make([]int64, 2)
	y := make([]int64, 2)
	r := make([]int64, 2)
	for i := int64(0); i < 2; i++ {
		scanner.Scan()
		x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		r[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(x, y, r)
}
