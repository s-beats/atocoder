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
	N, K := readIntSlice2()
	A := make([]int, N+1)
	for i := 0; i < K; i++ {
		l := readIntSlice()
		A[l[0]] = l[1]
	}

	dp := make([][][]int, N+2)
	for i := 0; i < N+2; i++ {
		dp[i] = make([][]int, 4)
		for j := 1; j < 4; j++ {
			dp[i][j] = make([]int, 4)
		}
	}

	// dp[3] までは先に求める
	bs := []int{1, 2, 3}
	if A[1] != 0 {
		bs = []int{A[1]}
	}
	for _, b := range bs {
		cs := []int{1, 2, 3}
		if A[2] != 0 {
			cs = []int{A[2]}
		}
		for _, c := range cs {
			dp[3][b][c] = 1
		}
	}

	// a-1日目までのパスタ
	// a-2日目に選んだのがb
	// a-1日目に選んだのがc
	for a := 3; a < N+1; a++ {
		for b := 1; b < 4; b++ {
			for c := 1; c < 4; c++ {
				kinds := []int{1, 2, 3}
				if A[a] != 0 {
					kinds = []int{A[a]}
				}

				for _, kind := range kinds {
					if kind == b && kind == c {
						continue
					}
					dp[a+1][c][kind] = (dp[a+1][c][kind] + dp[a][b][c]) % 10000
				}
			}
		}
	}

	ans := 0
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			ans = (ans + dp[N+1][i][j]) % 10000
		}
	}
	Println(ans)
}
