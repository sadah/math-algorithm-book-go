package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(L int64, R int64) {
	primes := make([]bool, R-L+1)
	for i := int64(0); i <= R-L; i++ {
		primes[i] = true
	}
	if L == 1 {
		primes[0] = false
	}

	for i := int64(2); i*i <= R; i++ {
		min := ((L + i - 1) / i) * i
		for j := min; j <= R; j += i {
			if j == i {
				continue
			}
			primes[j-L] = false
		}
	}

	ans := 0
	for _, v := range primes {
		if v {
			ans++
		}
	}
	fmt.Println(ans)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var L int64
	scanner.Scan()
	L, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var R int64
	scanner.Scan()
	R, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(L, R)
}
