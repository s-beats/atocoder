package main

import (
	. "fmt"
	"sort"
)

func main() {
	var N int
	Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		Scan(&A[i])
	}

	// 2,3のそれぞれの次数を求める
	two := make([]int, N)
	three := make([]int, N)
	prev := 1
	for i, v := range A {
		for v%2 == 0 {
			v /= 2
			two[i]++
		}
		for v%3 == 0 {
			v /= 3
			three[i]++
		}

		if i > 0 {
			if v != prev {
				Println(-1)
				return
			}
		}
		prev = v
	}

	// それぞれのリストの最小値と、各要素の差分を足す無名関数
	ans := 0
	fn := func(is []int) {
		sort.Slice(is, func(i, j int) bool { return is[i] < is[j] })
		for _, v := range is {
			ans += v - is[0]
		}
	}

	fn(two)
	fn(three)

	Println(ans)
}
