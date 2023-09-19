package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(N int64, X int64, A []int64) {
	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
	ret := binarySearch(A, X)
	if ret == -1 {
		fmt.Println(NO)
	} else {
		fmt.Println(YES)
	}
}

func binarySearch(nums []int64, target int64) (index int) {
	start := 0
	end := len(nums) - 1
	for {
		if end < start {
			return -1
		}
		index = (start + end) / 2

		if nums[index] == target {
			return index
		}
		if nums[index] < target {
			start = index + 1
		} else {
			end = index - 1
		}
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
	var X int64
	scanner.Scan()
	X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, X, A)
}
