package main

import "fmt"

// 配列の要素を添字として、値にはその要素の情報を保持する
// 言い換えると、配列の各要素の情報を保持する配列
// cf. https://terumiyake.hatenablog.com/entry/2020/05/19/234416

func main() {
	ss := []string{"a", "b", "a", "e", "e"}

	// 各要素の出現回数を保持する
	bucket1 := make(map[string]int)
	for _, s := range ss {
		bucket1[s]++
	}
	fmt.Println(bucket1)

	// 各要素の出現位置を保持する
	bucket2 := make(map[string][]int)
	for i, s := range ss {
		bucket2[s] = append(bucket2[s], i)
	}
	fmt.Println(bucket2)

	// 各要素の出現フラグを保持する
	bucket3 := make(map[string]bool)
	for _, v := range []string{"a", "b", "c", "d", "e"} {
		// falseで初期化
		bucket3[v] = false
	}
	for _, s := range ss {
		bucket3[s] = true
	}
	fmt.Println(bucket3)
}
