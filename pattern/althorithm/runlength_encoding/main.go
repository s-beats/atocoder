package main

import "fmt"

// 列の要素を区間分割する処理の1つとしてランレングス圧縮がある。
// ランレングス圧縮とは、連続する要素を1つの要素としてまとめる処理である。
// 本サンプルでは連続する要素と連続数のペアの二次元配列を生成する。

func main() {
	a := []string{"A", "A", "B", "C", "C", "A", "A", "D", "D", "D", "D"}
	var ans [][]any
	for i := 0; i < len(a); {
		v := a[i]
		j := i + 1
		count := 1
		for ; j < len(a) && a[j] == v; j++ {
			count++
		}
		ans = append(ans, []any{v, count})
		i = j
	}
	fmt.Println(a, ans)
}
