package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(H int64, W int64, A [][]int64) {
	rowSum := make([]int64, H)
	colSum := make([]int64, W)

	// Calculate row sum
	for i := int64(0); i < H; i++ {
		sum := int64(0)
		for j := int64(0); j < W; j++ {
			sum += A[i][j]
		}
		rowSum[i] = sum
	}

	// Calculate column sum
	for j := int64(0); j < W; j++ {
		sum := int64(0)
		for i := int64(0); i < H; i++ {
			sum += A[i][j]
		}
		colSum[j] = sum
	}

	for i := int64(0); i < H; i++ {
		ans := make([]string, W)
		for j := int64(0); j < W; j++ {
			ans[j] = strconv.Itoa(int(rowSum[i] + colSum[j] - A[i][j]))
		}
		fmt.Println(strings.Join(ans, " "))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var H int64
	scanner.Scan()
	H, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var W int64
	scanner.Scan()
	W, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([][]int64, H)
	for i := int64(0); i < H; i++ {
		A[i] = make([]int64, W)
	}
	for i := int64(0); i < H; i++ {
		for j := int64(0); j < W; j++ {
			scanner.Scan()
			A[i][j], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		}
	}
	solve(H, W, A)
}
