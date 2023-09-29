package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(N int64) {
	A := make([][]int64, 2)
	for i := 0; i < 2; i++ {
		A[i] = make([]int64, 2)
	}
	A[0][0] = 2
	A[0][1] = 1
	A[1][0] = 1
	A[1][1] = 0

	B := modPowerMatrix(A, N-1, MOD)
	ans := (B[1][0] + B[1][1]) % MOD
	fmt.Println(ans)
}

// A と B は n×n 行列
// A と B の積を m で割った余りを返す
func modMultiMatrix(A, B [][]int64, m int64) [][]int64 {
	n := len(A)
	C := make([][]int64, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int64, n)
	}
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			for j := 0; j < n; j++ {
				C[i][j] = (C[i][j] + A[i][k]*B[k][j]) % m
			}
		}
	}
	return C
}

// A は n×n 行列
// A の n 乗を m で割った余りを返す
func modPowerMatrix(A [][]int64, n, m int64) [][]int64 {
	if n == 0 {
		B := make([][]int64, len(A))
		for i := 0; i < len(A); i++ {
			B[i] = make([]int64, len(A))
			B[i][i] = 1
		}
		return B
	}
	B := modPowerMatrix(A, n/2, m)
	// B = A^(n/2) が求まったので、B^2 を計算する
	B = modMultiMatrix(B, B, m)
	// n が奇数の場合、B = B^2 * A を計算する
	if n%2 == 1 {
		B = modMultiMatrix(B, A, m)
	}
	return B
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
