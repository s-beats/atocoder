package main

import (
	"math"
	"sort"
)

func mod(n1, n2 int) int {
	res := (n1 + n2) % n2
	if res < 0 {
		res += n2
	}
	return res
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func popBack(l *[]interface{}) interface{} {
	e := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return e
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
}

func isOdd(n int) bool {
	return n&1 == 1
}

func isEven(n int) bool {
	return n&1 == 0
}

func perviousPermuutation(src []int) []int {
	size := len(src)
	dst := make([]int, size)
	copy(dst, src)

	// 入力の境目(大きい範囲の末尾)を探す
	boaderIdx := 0
	boader := -1
	for i := size - 1; i > 0; i-- {
		if src[i-1] > src[i] {
			boaderIdx = i - 1
			boader = src[i-1]
			break
		}
	}

	// 境目以降を降順
	sort.Slice(dst[boaderIdx:], func(i, j int) bool {
		return dst[i+boaderIdx] > dst[j+boaderIdx]
	})

	// 境目よりひとつ小さい値を探す
	for i := boaderIdx; i < size; i++ {
		if dst[i] < boader {
			dst[boaderIdx], dst[i] = dst[i], dst[boaderIdx]
			break
		}
	}

	// 境目+1以降を降順
	sort.Slice(dst[boaderIdx+1:], func(i, j int) bool {
		return dst[i+boaderIdx+1] > dst[j+boaderIdx+1]
	})

	return dst
}
