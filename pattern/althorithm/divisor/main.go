package main

import (
	"bytes"
	"os"
	"strconv"
)

func main() {
	// 約数の列挙
	n := 36
	ans := make([]int, n+1)
	for i := 1; i*i <= n; i++ {
		if ans[i] != 0 {
			continue
		}

		if n%i == 0 {
			ans[i] = 1
			ans[n/i] = 1
		} else {
			ans[i] = -1
			ans[n/i] = -1
		}
	}

	buf := bytes.Buffer{}
	for i, v := range ans {
		if v == 1 {
			_, _ = buf.WriteString(strconv.Itoa(i) + "\n")
		}
	}
	_, _ = buf.WriteTo(os.Stdout)
}
