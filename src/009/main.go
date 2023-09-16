package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(N int64, S int64, A []int64) {
	for bits := 0; bits < (1 << N); bits++ {
		sum := int64(0)
		for i := int64(0); i < N; i++ {
			if (bits>>uint64(i))&1 == 1 {
				sum += A[i]
			}
		}
		if sum == S {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
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
	var S int64
	scanner.Scan()
	S, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, S, A)
}
