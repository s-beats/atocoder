package main

import "fmt"

func main() {
	// N個の要素から0~N個選ぶパターンを列挙して調べる
	// 例えばN=3のとき、2^3=8パターン

	N := [3]int{1, 2, 3}
	W := 3
	Ans := [][]int{}

	// 2^3=8通りのbit列を全探索
	// 1<<3 - 1 = 7
	// 0 ~ 7
	for i := 0; i < 1<<len(N); i++ {
		sum := 0
		pattern := []int{}

		for j := 0; j < len(N); j++ {
			// 1をjビット左シフトして、1をAND演算
			// (e.g) 010 AND 001
			on := 1 << j & i
			if on > 0 {
				// jビット目が1であれば、j番目の要素を加算
				sum += N[j]
				pattern = append(pattern, N[j])
			}
		}

		if sum == W {
			Ans = append(Ans, pattern)
		}
	}

	for _, v := range Ans {
		fmt.Println(v)
	}

	// Output:
	// [1 2]
	// [3]
}
