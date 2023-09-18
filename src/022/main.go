package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, A []int64) {
	ans := int64(0)
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	m := make(map[int64]int64)
	for _, a := range A {
		m[a] += 1
	}

	keys := []int64{}
	for k := range m {
		keys = append(keys, int64(k))
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, k := range keys {
		if k > 100000-k {
			break
		}
		if c, ok := m[int64(100000-k)]; ok {
			if k == 100000-k {
				if c >= 2 {
					ans += c * (c - 1) / 2
				}
			} else {
				ans += m[int64(k)] * c
			}
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
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A)
}
