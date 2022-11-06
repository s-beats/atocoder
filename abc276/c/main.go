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

func perviousPermuutation(src []int) []int {
	size := len(src)
	dst := make([]int, size)
	copy(dst, src)

	// 入力の境目(大きい範囲の末尾)を探す
	boaderIdx := -1
	boader := -1
	for i := size - 1; i > 0; i-- {
		if src[i-1] > src[i] {
			boaderIdx = i - 1
			boader = src[i-1]
			break
		}
	}

	// 入力よりが一番前の場合はそのまま返す
	if boaderIdx == -1 {
		return dst
	}

	// 境目以降を降順
	sort.Slice(dst[boaderIdx:], func(i, j int) bool {
		return dst[i+boaderIdx] > dst[j+boaderIdx]
	})

	// 境目よりひとつ小さい値を探す
	for i := boaderIdx; i < size; i++ {
		if dst[i] < boader {
			dst[boaderIdx], dst[i] = dst[i], dst[boaderIdx]
			break
		}
	}

	// 境目+1以降を降順
	sort.Slice(dst[boaderIdx+1:], func(i, j int) bool {
		return dst[i+boaderIdx+1] > dst[j+boaderIdx+1]
	})

	return dst
}

func main() {
	_ = toInt(readline())
	P := readIntSlice()

	var ans string
	for _, v := range perviousPermuutation(P) {
		ans += strconv.Itoa(v) + " "
	}
	Println(strings.TrimRight(ans, " "))
}
