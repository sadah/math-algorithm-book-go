package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64) {
	ret := primeFactorize(N)

	var keys []int64
	for k := range ret {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, k := range keys {
		v := ret[k]
		if v != int64(0) {
			for j := int64(1); j <= v; j++ {
				fmt.Print(k, " ")
			}
		}
	}
	fmt.Println()
}

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
