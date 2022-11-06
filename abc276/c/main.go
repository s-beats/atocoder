package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
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

func main() {
	N := toInt(readline())
	P := readIntSlice()

	// 入力の境目(大きい範囲の末尾)を探す
	boaderIdx := 0
	boader := -1
	for i := N - 1; i > 0; i-- {
		if P[i-1] > P[i] {
			boaderIdx = i - 1
			boader = P[i-1]
			break
		}
	}

	// 境目以降を降順
	sort.Slice(P[boaderIdx:], func(i, j int) bool {
		return P[i+boaderIdx] > P[j+boaderIdx]
	})

	// 境目よりひとつ小さい値を探す
	for i := boaderIdx; i < N; i++ {
		if P[i] < boader {
			P[boaderIdx], P[i] = P[i], P[boaderIdx]
			break
		}
	}

	// 境目+1以降を降順
	sort.Slice(P[boaderIdx+1:], func(i, j int) bool {
		return P[i+boaderIdx+1] > P[j+boaderIdx+1]
	})

	var ans string
	for i := 0; i < N; i++ {
		ans += strconv.Itoa(P[i]) + " "
	}
	Println(strings.TrimRight(ans, " "))
}
