package lib

import (
	"math/big"
	"sort"
)

// isPrime returns true if n is prime.
func isPrime(n int64) bool {
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// divisors returns all divisors of N.
func divisors(N int64) []int64 {
	var ret []int64
	for i := int64(1); i*i <= N; i++ {
		if N%i == 0 {
			ret = append(ret, i)
			if i*i != N {
				ret = append(ret, N/i)
			}
		}
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})
	return ret
}

// primeFactorize returns prime factorization of N.
func primeFactorize(N int64) map[int64]int64 {
	ret := make(map[int64]int64)
	for i := int64(2); i*i <= N; i++ {
		for N%i == 0 {
			ret[i]++
			N /= i
		}
	}
	if N != 1 {
		ret[N] = 1
	}
	return ret
}

// gcd returns the greatest common divisor of a and b.
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// lcm returns the least common multiple of a and b.
func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}

// factorial returns n!.
func factorial(n int64) int64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// combination returns nCr.
func combination(n, r int64) int64 {
	if n < r {
		return 0
	}
	return factorial(n) / (factorial(r) * factorial(n-r))
}

var fact []*big.Int

// modCombination returns nCr % MOD.
func modCombination(n, r, mod int64) int64 {
	if len(fact) != int(n+1) {
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

// modCombination returns nCr % MOD.
func modCombinationNaive(n, r, mod int64) int64 {
	numerator, denominator := int64(1), int64(1)
	// n C r = n! / (r! * (n-r)!)
	// n! % MOD
	for i := int64(1); i <= n; i++ {
		numerator = (numerator * int64(i)) % mod
	}
	// r! % MOD
	for i := int64(1); i <= r; i++ {
		denominator = (denominator * int64(i)) % mod
	}

	// (n-r)! % MOD
	for i := int64(1); i <= n-r; i++ {
		denominator = (denominator * int64(i)) % mod
	}

	return modDiv(numerator, denominator, mod)
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
