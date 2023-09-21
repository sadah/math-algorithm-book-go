package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(x []int64, y []int64) {
	x1, x2, x3, x4 := x[0], x[1], x[2], x[3]
	y1, y2, y3, y4 := y[0], y[1], y[2], y[3]

	// 2つのベクトルの外積を計算
	vpABCD := vectorProduct(x2-x1, y2-y1, x3-x1, y3-y1)
	vpABDB := vectorProduct(x2-x1, y2-y1, x4-x1, y4-y1)
	vpCDAD := vectorProduct(x4-x3, y4-y3, x1-x3, y1-y3)
	vpCDBD := vectorProduct(x4-x3, y4-y3, x2-x3, y2-y3)

	// 両方のベクトルがゼロベクトルの場合、2つの線分は同じ直線上にある
	if vpABCD == 0 && vpABDB == 0 && vpCDAD == 0 && vpCDBD == 0 {
		// 線分をソートし、最大と最小の座標を取得
		x1, y1, x2, y2 = sortPoints(x1, y1, x2, y2)
		x3, y3, x4, y4 = sortPoints(x3, y3, x4, y4)
		_, _, maxX, maxY := sortPoints(x1, y1, x3, y3)
		minX, minY, _, _ := sortPoints(x2, y2, x4, y4)

		// 交差しているかどうかを判定
		if maxX < minX || (maxX == minX && maxY < minY || (maxX == minX && maxY == minY)) {
			fmt.Println(YES)
		} else {
			fmt.Println(NO)
		}
		return
	}

	// ベクトルABとCDがそれぞれ交差しているかどうかを判定
	// 2つの線分ABとCDが交差しているかどうかを判定するには、それぞれの線分をベクトルで表現し、そのベクトル同士の外積を計算
	// 交差がある場合、外積ABxCDと外積CDxABの符号が異なる
	// 外積がゼロの場合、線分ABとCDは平行であるか同一直線上にあるため、交差はない
	crossedAB, crossedCD := false, false
	if vpABCD >= 0 && vpABDB <= 0 {
		crossedAB = true
	}
	if vpABCD <= 0 && vpABDB >= 0 {
		crossedAB = true
	}
	if vpCDAD >= 0 && vpCDBD <= 0 {
		crossedCD = true
	}
	if vpCDAD <= 0 && vpCDBD >= 0 {
		crossedCD = true
	}

	if crossedAB && crossedCD {
		fmt.Println(YES)
	} else {
		fmt.Println(NO)
	}
}

// 2つの座標をソートする
func sortPoints(x1, y1, x2, y2 int64) (int64, int64, int64, int64) {
	if x1 > x2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	} else if x1 == x2 {
		if y1 > y2 {
			y1, y2 = y2, y1
		}
	}
	return x1, y1, x2, y2
}

// ベクトルAとベクトルBの外積（cross product）は、次のように計算される
// 外積 AB = Ax * By - Ay * Bx
// AxとAyはベクトルAのx成分とy成分であり、BxとByはベクトルBのx成分とy成分
// 外積の結果が正の値の場合、ベクトルAとベクトルBは時計回りに向かっていることを示し、負の値の場合は反時計回りに向かっていることを示す
// 外積がゼロの場合、ベクトルAとベクトルBは平行であり、同一直線上にあることを示す
func vectorProduct(x1 int64, y1 int64, x2 int64, y2 int64) int64 {
	return x1*y2 - x2*y1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	x := make([]int64, 4)
	y := make([]int64, 4)
	for i := int64(0); i < 4; i++ {
		scanner.Scan()
		x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(x, y)
}
