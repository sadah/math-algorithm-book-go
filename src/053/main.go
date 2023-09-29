package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(N int64) {
	v := modpow(4, N+1, MOD) - 1
	ans := division(v, 3, MOD)
	fmt.Println(ans)
}

func modpow(a, b, m int64) int64 {
	// 繰り返し二乗法（p は a^1, a^2, a^4, a^8, ... といった値をとる）
	p := a
	answer := int64(1)
	for i := 0; i < 60; i++ {
		if (b & (1 << i)) != 0 {
			answer = (answer * p) % m
		}
		p = (p * p) % m
	}
	return answer
}

func division(a, b, m int64) int64 {
	// division(a, b, m) は a÷b mod m を返す関数
	// a÷b = a×b^-1 となる
	// b^-1 は b の逆元である
	// b^-1 は b^(m-2) mod m で求められる
	// これはフェルマーの小定理より成り立つ
	return (a * modpow(b, m-2, m)) % m
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
