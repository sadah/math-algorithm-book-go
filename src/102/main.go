package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, c []string) {

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
    c := make([]string, 1)
    for i := int64(0); i < 1; i++ {
        scanner.Scan()
        c[i] = scanner.Text()
    }
	solve(N, c)
}
