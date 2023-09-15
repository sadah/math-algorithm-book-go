package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func solve(A []int64) {
	sum := int64(0)
	for _, x := range A {
		sum += x
	}
	fmt.Println(sum)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	A := make([]int64, 3)
    for i := int64(0); i < 3; i++ {
        scanner.Scan()
        A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(A)
}
