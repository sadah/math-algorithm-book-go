package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, schedules []schedule) {
	sort.Slice(schedules, func(i, j int) bool {
		return schedules[i].end < schedules[j].end
	})

	ans := int64(1)
	cur := schedules[0].end

	for i := int64(0); i < N; i++ {
		if schedules[i].start < cur {
			continue
		}
		cur = schedules[i].end
		ans++
	}
	fmt.Println(ans)
}

type schedule struct {
	start int64
	end   int64
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
	schedules := make([]schedule, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		schedules[i].start, _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		schedules[i].end, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	}
	solve(N, schedules)
}
