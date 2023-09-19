package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, A []int64) {
	for _, a := range mergeSort(A) {
		fmt.Print(a, " ")
	}
	fmt.Println()
}

// merge sort
func mergeSort(A []int64) []int64 {
	if len(A) <= 1 {
		return A
	}
	mid := len(A) / 2
	left := mergeSort(A[:mid])
	right := mergeSort(A[mid:])
	return merge(left, right)
}

// merge two sorted arrays
func merge(left, right []int64) []int64 {
	ret := []int64{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}
	return append(ret, append(left, right...)...)
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
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A)
}
