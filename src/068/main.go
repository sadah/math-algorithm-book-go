package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, K int64, V []int64) {
	ans := int64(0)

	// bit search
	for i := int64(1); i < (1 << K); i++ {
		cnt := int64(0)
		l := int64(1)
		for j := int64(0); j < K; j++ {
			if i&(1<<j) != 0 {
				cnt++
				l = lcm(l, V[j])
			}
		}
		n := N / l
		if cnt%2 == 1 {
			ans += n
		} else {
			ans -= n
		}
	}
	fmt.Println(ans)
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	V := make([]int64, K)
	for i := int64(0); i < K; i++ {
		scanner.Scan()
		V[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, K, V)
}
