package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, X int64) {
	ans := 0
	for a := int64(1); a <= N; a++ {
		for b := a + int64(1); b <= N; b++ {
			for c := b + int64(1); c <= N; c++ {
				if a+b+c == X {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
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
	solve(N, X)
}
