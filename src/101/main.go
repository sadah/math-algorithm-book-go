package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(N int64) {
	for i := int64(1); i <= N; i++ {
		ans := int64(0)
		for j := int64(1); j <= (N+i-1)/i; j++ {
			ans += modCombination(N-(i-1)*(j-1), j, MOD)
			ans %= MOD
		}
		fmt.Println(ans)
	}
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

var fact []*big.Int

// modCombination returns nCr % MOD.
func modCombination(n, r, mod int64) int64 {
	if len(fact) == 0 {
		fact = make([]*big.Int, n+1)
		fact[0] = big.NewInt(1)
		for i := int64(1); i <= n; i++ {
			fact[i] = new(big.Int)
			fact[i].Set(fact[i-1])
			fact[i].Mul(fact[i], big.NewInt(i))
			fact[i].Mod(fact[i], big.NewInt(mod))
		}
	}
	numerator := new(big.Int)
	numerator.Set(fact[n])
	denominator := new(big.Int)
	denominator.Mul(fact[r], fact[n-r])
	denominator.Mod(denominator, big.NewInt(mod))
	return modDiv(numerator.Int64(), denominator.Int64(), mod)
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
