package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(N int64, M int64, A []int64, B []int64, C []int64) {
	// 隣接リストの作成
	G := make([][]Edge, N+1)
	for i := int64(1); i <= N; i++ {
		G[i] = make([]Edge, 0)
	}
	for i := int64(1); i <= M; i++ {
		G[A[i]] = append(G[A[i]], Edge{B[i], C[i]})
		G[B[i]] = append(G[B[i]], Edge{A[i], C[i]})
	}

	dist := dijkstra(N, M, G)
	if dist[N] != math.MaxInt64 {
		fmt.Println(dist[N])
	} else {
		fmt.Println(-1)
	}
}

func dijkstra(N int64, M int64, G [][]Edge) []int64 {
	// ダイクストラ法：配列の初期化など
	dist := make([]int64, N+1)
	used := make([]bool, N+1)
	for i := int64(1); i <= N; i++ {
		dist[i] = math.MaxInt64
		used[i] = false
	}

	Q := make(PriorityQueue, 0)
	dist[1] = 0
	heap.Init(&Q)
	heap.Push(&Q, Pair{0, 1})

	// ダイクストラ法：優先度付きキューの更新
	for Q.Len() > 0 {
		item := heap.Pop(&Q).(Pair)
		pos := int(item.second)
		if used[pos] {
			continue
		}
		used[pos] = true
		for _, i := range G[pos] {
			to := i.to
			cost := dist[pos] + i.cost
			if dist[to] > cost {
				dist[to] = cost
				heap.Push(&Q, Pair{dist[to], int64(to)})
			}
		}
	}
	return dist
}

type Edge struct {
	to   int64
	cost int64
}

type Pair struct {
	first, second int64
}

type PriorityQueue []Pair

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].first < pq[j].first }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Pair)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
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
	var M int64
	scanner.Scan()
	M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, M+1)
	B := make([]int64, M+1)
	C := make([]int64, M+1)
	for i := int64(1); i <= M; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		C[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, M, A, B, C)
}
