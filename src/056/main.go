package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(N int64) {
	A := make([][]int64, 3)
	for i := 0; i < 3; i++ {
		A[i] = make([]int64, 3)
	}
	A[0][0] = 1
	A[0][1] = 1
	A[0][2] = 1
	A[1][0] = 1
	A[1][1] = 0
	A[1][2] = 0
	A[2][0] = 0
	A[2][1] = 1
	A[2][2] = 0

	B := modPowerMatrix(A, N-1, MOD)
	ans := B[0][0] % MOD

	fmt.Println(ans)
}

// A と B は n×n 行列
// A と B の積を m で割った余りを返す
func modMultiMatrix(A, B [][]int64, m int64) [][]int64 {
	// 行列 A の次元 (n x n)
	n := len(A)
	// 結果の行列 C を生成
	C := make([][]int64, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int64, n)
	}
	// 行列の掛け算を実行
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			for j := 0; j < n; j++ {
				// C[i][j] = (A[i][0] * B[0][j] + A[i][1] * B[1][j] + ... + A[i][k] * B[k][j]) % m
				C[i][j] = (C[i][j] + A[i][k]*B[k][j]) % m
			}
		}
	}
	return C
}

// A は n×n の行列で A の n 乗を m で割った余りを返す
func modPowerMatrix(A [][]int64, n, m int64) [][]int64 {
	// n が 0 の場合、単位行列（対角成分がすべて 1 でそれ以外が 0 の行列）を生成して返す
	// これは n のゼロ乗を表す
	if n == 0 {
		B := make([][]int64, len(A))
		for i := 0; i < len(A); i++ {
			B[i] = make([]int64, len(A))
			B[i][i] = 1
		}
		return B
	}
	// n が 0 でない場合、A の n 乗を計算するために再帰的に関数を呼び出す
	// n を半分に割った指数で行列 A のべき乗を計算する
	// これを行うために、modPowerMatrix 関数を再帰的に呼び出す
	B := modPowerMatrix(A, n/2, m)
	// B = A^(n/2) が求まったので、B^2 を計算する
	// modMultiMatrix 関数で、2つの行列を掛け算し、結果を m で割った余りを計算する
	B = modMultiMatrix(B, B, m)
	// n が奇数の場合、B = B^2 * A を計算する
	// これにより指数 n の行列 A のべき乗を求る
	if n%2 == 1 {
		B = modMultiMatrix(B, A, m)
	}
	// 結果の行列 B は行列 A の n 乗を m で割った余りを表す
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
