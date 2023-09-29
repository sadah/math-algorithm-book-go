package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(X int64, Y int64) {
	if (2*Y-X) < 0 || (2*X-Y) < 0 {
		fmt.Println(0)
	} else if (2*Y-X)%3 != 0 || (2*X-Y)%3 != 0 {
		fmt.Println(0)
	} else {
		a := (2*Y - X) / 3
		b := (2*X - Y) / 3
		numerator, denominator := int64(1), int64(1)

		// x+y C x = (x+y)! / (x! * y!)
		// (x+y)! % MOD
		for i := int64(1); i <= a+b; i++ {
			numerator = (numerator * int64(i)) % MOD
		}

		// x! % MOD
		for i := int64(1); i <= a; i++ {
			denominator = (denominator * int64(i)) % MOD
		}

		// y! % MOD
		for i := int64(1); i <= b; i++ {
			denominator = (denominator * int64(i)) % MOD
		}

		ans := division(numerator, denominator, MOD)
		fmt.Println(ans)
	}
}

func modpow(a, b, m int64) int64 {
	// 繰り返し二乗法（p は a^1, a^2, a^4, a^8, ... といった値をとる）
	p := a
	answer := int64(1)
	for i := 0; i < 30; i++ {
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
	var X int64
	scanner.Scan()
	X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var Y int64
	scanner.Scan()
	Y, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(X, Y)
}
