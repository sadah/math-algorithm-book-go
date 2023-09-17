package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64) {
	divs := divisors(N)
	for _, d := range divs {
		fmt.Println(d)
	}
}

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
