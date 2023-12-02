package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(Q int64, X []float64, Y []float64, Z []float64, T []int64) {
	for i := int64(0); i < Q; i++ {
		var a [3][3]float64
		a[0][0] = 1.0 - X[i]
		a[2][0] = X[i]
		a[1][1] = 1.0 - Y[i]
		a[0][1] = Y[i]
		a[2][2] = 1.0 - Z[i]
		a[1][2] = Z[i]

		b := power(a, T[i])
		ansA := b[0][0] + b[0][1] + b[0][2]
		ansB := b[1][0] + b[1][1] + b[1][2]
		ansC := b[2][0] + b[2][1] + b[2][2]
		fmt.Println(ansA, ansB, ansC)
	}
}

func power(a [3][3]float64, n int64) [3][3]float64 {
	p := a
	var q [3][3]float64
	flag := false
	for i := 0; i < 30; i++ {
		if (n & (1 << i)) != 0 {
			if !flag {
				q = p
				flag = true
			} else {
				q = multiplication(q, p)
			}
		}
		p = multiplication(p, p)
	}
	return q
}

func multiplication(a, b [3][3]float64) [3][3]float64 {
	var c [3][3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var Q int64
	scanner.Scan()
	Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	X := make([]float64, Q)
	Y := make([]float64, Q)
	Z := make([]float64, Q)
	T := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		X[i], _ = strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		Y[i], _ = strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		Z[i], _ = strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		T[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(Q, X, Y, Z, T)
}
