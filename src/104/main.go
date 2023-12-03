package main

import "fmt"

func main() {
	n, m := 9, 12
	g := [12][2]int{
		{1, 2},
		{2, 3},
		{4, 5},
		{5, 6},
		{7, 8},
		{8, 9},
		{1, 4},
		{4, 7},
		{2, 5},
		{5, 8},
		{3, 6},
		{6, 9},
	}

	fmt.Println(n, m)
	for _, p := range g {
		fmt.Println(p[0], p[1])
	}
}
