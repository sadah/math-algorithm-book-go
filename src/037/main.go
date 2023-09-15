package main

import (
	"bufio"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(x []int64, y []int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	x := make([]int64, 4)
    y := make([]int64, 4)
    for i := int64(0); i < 4; i++ {
        scanner.Scan()
        x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(x, y)
}
