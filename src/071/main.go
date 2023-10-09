package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, a []int64, b []int64, c []int64) {
	ans := float64(0)
	for i := int64(0); i < N; i++ {
		for j := int64(0); j < N; j++ {
			if a[i]*b[j] == a[j]*b[i] {
				continue
			}
			px := float64(c[i]*b[j]-c[j]*b[i]) / float64(a[i]*b[j]-a[j]*b[i])
			py := float64(c[i]*a[j]-c[j]*a[i]) / float64(b[i]*a[j]-b[j]*a[i])
			ret := check(N, a, b, c, px, py)
			if ret {
				ans = max(ans, px+py)
			}
		}
	}
	fmt.Println(ans)
}

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func check(N int64, a, b, c []int64, x, y float64) bool {
	for i := int64(0); i < N; i++ {
		if float64(a[i])*x+float64(b[i])*y > float64(c[i]) {
			return false
		}
	}
	return true
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
	a := make([]int64, N)
	b := make([]int64, N)
	c := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		b[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		c[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, a, b, c)
}
