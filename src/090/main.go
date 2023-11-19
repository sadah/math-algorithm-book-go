package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var fmCand map[int64]struct{}

func solve(N int64, B int64) {
	fmCand = make(map[int64]struct{})
	ans := int64(0)
	f(0, 0)
	for fm := range fmCand {
		m := fm + B
		prodM := product(m)
		if m-prodM == B && m <= N {
			ans += 1
		}
	}
	fmt.Println(ans)
}

func product(m int64) int64 {
	if m == 0 {
		return 0
	} else {
		ans := int64(1)
		for m >= 1 {
			ans *= m % 10
			m /= 10
		}
		return ans
	}
}

func f(digit int, m int64) {
	if digit == 11 {
		fmCand[product(m)] = struct{}{}
		return
	}
	minValue := m % 10
	for i := minValue; i <= 9; i++ {
		f(digit+1, 10*m+i)
	}
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
	var B int64
	scanner.Scan()
	B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N, B)
}
