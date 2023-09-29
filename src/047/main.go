package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(N int64, M int64, A []int64, B []int64) {
	adj := adjacencyList(N, M, A, B)
	if isBipartiteGraph(N, adj) {
		fmt.Println(YES)
	} else {
		fmt.Println(NO)
	}
}

// 与えられたグラフの隣接リストを生成する
// N はノード数、M は辺数、A と B は辺の始点ノードと終点ノードを示すスライス
// 隣接リストは、各ノードから直接アクセスできる隣接ノードのリストを表す
func adjacencyList(N int64, M int64, A []int64, B []int64) (adj [][]int64) {
	adj = make([][]int64, N)
	for i := int64(0); i < M; i++ {
		adj[A[i]-1] = append(adj[A[i]-1], B[i]-1)
		adj[B[i]-1] = append(adj[B[i]-1], A[i]-1)
	}
	return
}

// 与えられたグラフが二部グラフであるかどうかを判定する
// グラフのノードに色を割り当てて、隣接ノードが異なる色で塗られているかどうかを確認する
// 隣接ノードが同じ色であれば、二部グラフではない
// 深さ優先探索（DFS）を使用して、ノードに色を塗りながら隣接ノードを探索する
func isBipartiteGraph(N int64, adj [][]int64) bool {
	// 0-indexed
	color := make([]int64, N)
	for i := int64(0); i < N; i++ {
		if color[i] == 0 {
			if !dfs(i, 1, color, adj) {
				return false
			}
		}
	}
	return true
}

func dfs(v, c int64, color []int64, adj [][]int64) bool {
	color[v] = c
	for _, u := range adj[v] {
		if color[u] == c {
			return false
		}
		if color[u] == 0 && !dfs(u, -c, color, adj) {
			return false
		}
	}
	return true
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
	A := make([]int64, M)
	B := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, M, A, B)
}
