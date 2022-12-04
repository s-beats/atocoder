package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	// 素因数分解
	n := 1000
	ans := make([][]int, 0, int(math.Sqrt(float64(n))))
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			exp := 0
			for n%i == 0 {
				n /= i
				exp++
			}
			ans = append(ans, []int{i, exp})
		}
	}
	if n != 1 {
		ans = append(ans, []int{n, 1})
	}

	buf := bytes.NewBuffer(nil)
	for _, v := range ans {
		_, _ = buf.WriteString(
			fmt.Sprintf("%d %d\n", v[0], v[1]),
		)
	}
	_, _ = buf.WriteTo(os.Stdout)
}
