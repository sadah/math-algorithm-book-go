package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dy = []int64{-1, 0, 1, 0}
var dx = []int64{0, 1, 0, -1}

type Point struct {
	y, x int64
}

func solve(R, C, sy, sx, gy, gx int64, c [][]string) {
	// 0-indexed
	sy--
	sx--
	gy--
	gx--

	visited := make([][]bool, R)
	for i := int64(0); i < R; i++ {
		visited[i] = make([]bool, C)
	}
	visited[sy][sx] = true

	m := map[Point]int64{}
	queue := []Point{{sy, sx}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			ny, nx := p.y+dy[i], p.x+dx[i]
			if c[ny][nx] == "." && !visited[ny][nx] {
				p := Point{ny, nx}
				queue = append(queue, p)
				visited[ny][nx] = true
				m[p] = m[Point{p.y - dy[i], p.x - dx[i]}] + 1
			}
		}
	}
	fmt.Println(m[Point{gy, gx}])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var R int64
	scanner.Scan()
	R, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	var C int64
	scanner.Scan()
	C, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	var sy int64
	scanner.Scan()
	sy, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	var sx int64
	scanner.Scan()
	sx, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	var gy int64
	scanner.Scan()
	gy, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	var gx int64
	scanner.Scan()
	gx, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	c := make([][]string, R)
	for i := int64(0); i < R; i++ {
		scanner.Scan()
		c[i] = strings.Split(scanner.Text(), "")
	}
	solve(R, C, sy, sx, gy, gx, c)
}
