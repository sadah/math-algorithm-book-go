package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(K int64) {
	fmt.Println(findMinimumSumDigits(K))
}

func findMinimumSumDigits(k int64) int64 {
	// 各値までの最短距離を格納するスライスを初期化
	dist := make([]int64, k)
	for i := int64(0); i < k; i++ {
		// 初期値を大きな値に設定
		dist[i] = 1 << 29
	}

	// BFSのためのキューを初期化
	queue := make([]int64, 0, k)
	dist[1] = 1
	queue = append(queue, 1)

	// BFSを実行
	for len(queue) > 0 {
		val := queue[0]
		queue = queue[1:]

		// valを10倍した値
		val10 := (val * 10) % k
		// 各桁の和が同じで、より短い距離が見つかれば更新
		if dist[val10] > dist[val] {
			dist[val10] = dist[val]
			queue = append([]int64{val10}, queue...)
		}

		// valに1を加えた値を計算
		val1 := (val + 1) % k
		// 各桁の和が異なり、より短い距離が見つかれば更新
		if dist[val1] > dist[val]+1 {
			dist[val1] = dist[val] + 1
			queue = append(queue, val1)
		}
	}

	// 各桁の和が最小の倍数がdist[0]に格納されている
	return dist[0]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(K)
}
