package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// 1行の文字列がデフォルトを超えることがあるのでサイズを明示的に指定
var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readline() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, err := rdr.ReadLine()
		if err != nil {
			panic(err)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func readIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, toInt(v))
	}
	return slice
}

func readIntSlice2() (a, b int) {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, toInt(v))
	}
	return slice[0], slice[1]
}

func readIntSlice3() (a, b, c int) {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, toInt(v))
	}
	return slice[0], slice[1], slice[2]
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func toStr(n int) string {
	return strconv.Itoa(n)
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func isOdd(n int) bool {
	return n&1 == 1
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 1文字
func isUpper(s string) bool {
	return unicode.IsUpper(rune(s[0]))
}

// 1文字
func isLower(s string) bool {
	return unicode.IsUpper(rune(s[0]))
}

var (
	dr = []int{1, -1, 0, 0}
	dc = []int{0, 0, 1, -1}
)

func main() {
	W := toInt(readline())
	N, K := readIntSlice2()
	A, B := make([]int, N+1), make([]int, N+1)
	for i := 0; i < N; i++ {
		l := readIntSlice()
		A[i+1] = l[0]
		B[i+1] = l[1]
	}

	dp := make([][][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([][]int, K+1)
		for j := 0; j < K+1; j++ {
			dp[i][j] = make([]int, W+1)
		}
	}

	ans := 0
	fn := func(i, j, k int) {
		defer func() {
			ans = max(ans, dp[i][j][k])
		}()

		// 追加無し
		dp[i][j][k] = dp[i-1][j][k]

		// 追加有り
		//
		// 使用するカード->i-1枚目まで
		// 使用するカード数->j-1枚目以下
		// 幅->k-A[i]
		// 上記条件から遷移
		//
		// 幅がkを超える場合は追加無し
		if A[i] > k {
			return
		}
		dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-1][k-A[i]]+B[i])
	}

	for i := 1; i < N+1; i++ {
		for j := 1; j < K+1; j++ {
			for k := 1; k < W+1; k++ {
				fn(i, j, k)
			}
		}
	}

	Println(ans)
}
