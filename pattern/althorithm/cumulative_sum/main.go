package main

import "fmt"

func main() {
	// 累積和
	// 一次元数列
	slice := []int{1, 0, 9, 8, 2, 3, 5}

	// 添字x~添字yまでの和を求める
	// 毎回ループ
	fnSlow := func(x, y int) int {
		sum := 0
		for i := x; i < y; i++ {
			sum += slice[i]
		}
		return sum
	}
	// 最悪計算量: O(N)
	// Nは数列の長さ
	// クエリが複数ある場合は、O(QN)
	// Qはクエリの数
	fmt.Println(fnSlow(0, 6)) // 28
	fmt.Println(fnSlow(1, 3)) // 17

	// 累積和を使う
	// 先に累積和を求めておく
	cumulativeSum := make([]int, len(slice)+1)
	cumulativeSum[0] = 0 // 先頭は必ず0、先頭からi個の要素を求めるため
	for i := 0; i < len(slice); i++ {
		cumulativeSum[i+1] = cumulativeSum[i] + slice[i]
	}
	// 最悪計算量: O(N)
	// Nは数列の長さ
	// クエリが複数ある場合は、O(Q+N)
	// 各クエリはO(1)で求められる
	fmt.Println(cumulativeSum[6] - cumulativeSum[0]) // 28
	fmt.Println(cumulativeSum[3] - cumulativeSum[1]) // 17
}
