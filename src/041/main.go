package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(T int64, N int64, L []int64, R []int64) {
	nums := make([]int64, T+2)
	for i := int64(0); i < N; i++ {
		nums[L[i]+1]++
		nums[R[i]+1]--
	}
	for i := int64(0); i < T; i++ {
		nums[i+1] += nums[i]
		fmt.Println(nums[i+1])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var T int64
	scanner.Scan()
	T, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	L := make([]int64, N)
	R := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		L[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		R[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(T, N, L, R)
}
