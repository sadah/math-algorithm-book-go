package lib

import "sort"

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

// combination returns nCr.
func combination(n, r int64) int64 {
	if n < r {
		return 0
	}
	if n == r {
		return 1
	}
	if r == 1 {
		return n
	}
	return combination(n-1, r) + combination(n-1, r-1)
}
