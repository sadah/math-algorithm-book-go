package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64) {
	if N%4 == 1 {
		fmt.Println(2)
	} else if N%4 == 2 {
		fmt.Println(4)
	} else if N%4 == 3 {
		fmt.Println(8)
	} else if N%4 == 0 {
		fmt.Println(6)
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
	solve(N)
}
