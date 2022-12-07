package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 1行の文字列がデフォルトを超えることがあるのでサイズを明示的に指定
var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readline() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			Println(e.Error())
			panic(e)
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

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func isOdd(n int) bool {
	return n&1 == 1
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

var (
	N, M, Q int
	rows    [][]int
	A       []int
)

func main() {
	l := readIntSlice()
	N, M, Q = l[0], l[1], l[2]
	rows = make([][]int, Q)
	for i := 0; i < Q; i++ {
		rows[i] = readIntSlice()
	}

	A = make([]int, 1, N)
	A[0] = 1
	Println(rec())
}

// 再帰
func rec() int {
	size := len(A)

	// 数列が完成したら最大値を求める
	if N == size {
		sum := 0
		for _, v := range rows {
			Abi := (A)[v[1]-1]
			Bai := (A)[v[0]-1]
			ci := v[2]
			di := v[3]
			if Abi-Bai == ci {
				sum += di
			}
		}
		return sum
	}

	// 数列を作りつつ最大値を求める
	// [1] -> [1 1] -> [1 1 1] -> [1 1 1 1] -> [1 1 1 2] -> [1 1 1 3] ... [1 1 1 4] -> [1 1 2 2] ...
	// e.g N=4であれば[1 1 1 1] ~ [4 4 4 4]まで
	ret := 0
	last := A[size-1]
	for ; last <= M; last++ {
		A = append(A, last)
		ret = max(ret, rec())
		A = (A)[:size]
	}

	return ret
}
