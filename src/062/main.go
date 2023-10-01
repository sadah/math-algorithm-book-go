package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, K int64, A []int64) {
	first, second := make([]int64, N), make([]int64, N)
	for i := int64(0); i < N; i++ {
		first[i], second[i] = -1, -1
	}

	cnt, cur := int64(0), int64(0)
	for {
		if first[cur] == -1 {
			first[cur] = cnt
		} else if second[cur] == -1 {
			second[cur] = cnt
		}

		if cnt == K {
			fmt.Println(cur + 1)
			return
		} else if second[cur] != -1 && (K-first[cur])%(second[cur]-first[cur]) == 0 {
			fmt.Println(cur + 1)
			return
		}
		cur = A[cur] - 1
		cnt += 1
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
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, K, A)
}
