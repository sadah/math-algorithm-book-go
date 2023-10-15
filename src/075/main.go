package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(N int64, A []int64) {
	ans := int64(0)
	for i := int64(1); i <= N; i++ {
		ans += A[i] * modCombination(N-1, i-1, MOD)
		ans %= MOD
	}
	fmt.Println(ans)
}

// a の n 乗を m で割った余りを返す関数
func modPower(a, n, m int64) int64 {
	// n が 0 の場合、1 を返す
	if n == 0 {
		return 1
	}
	// n が偶数の場合、a^n = (a^(n/2))^2 を利用して計算
	if n%2 == 0 {
		halfPower := modPower(a, n/2, m)
		return (halfPower * halfPower) % m
	}
	// n が奇数の場合、x^n = x * (x^(n-1)) を利用して計算
	return (a * modPower(a, n-1, m)) % m
}

// modDiv(a, b, m) は a÷b mod m を返す関数
// a÷b = a×b^-1 となる
// b^-1 は b の逆元である
// b^-1 は b^(m-2) mod m で求められる
// これはフェルマーの小定理より成り立つ
func modDiv(a, b, m int64) int64 {
	return (a * modPower(b, m-2, m)) % m
}

var fact []int64

// modCombination returns nCr % MOD.
func modCombination(n, r, mod int64) int64 {
	if len(fact) == 0 {
		fact = make([]int64, n+1)
		fact[0] = 1
		for i := int64(1); i <= n; i++ {
			fact[i] = (fact[i-1] * i) % mod
		}
	}
	return modDiv(fact[n], fact[r]*fact[n-r], mod)
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
