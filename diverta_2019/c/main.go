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

// ペアリング
// 組み合わせと上界をうまく求める
func main() {
	N := toInt(readline())
	NS := make([]string, N)

	// `AB`を含む文字列の数
	includeAB := 0
	isIncludeAB := func(s string) bool {
		return strings.Contains(s, "AB")
	}
	countAB := func(s string) int {
		return strings.Count(s, "AB")
	}

	// `B`で始まり`A`で終わる文字列の数
	startBendA := 0
	isStartBendA := func(s string) bool {
		return strings.HasPrefix(s, "B") && strings.HasSuffix(s, "A")
	}

	// `B`で始まる文字列の数
	startB := 0
	isStartB := func(s string) bool {
		return strings.HasPrefix(s, "B")
	}

	// `A`で終わる文字列の数
	endA := 0
	isEndA := func(s string) bool {
		return strings.HasSuffix(s, "A")
	}

	for i := 0; i < N; i++ {
		NS[i] = readline()
		s := NS[i]

		if isIncludeAB(s) {
			includeAB += countAB(s)
		}

		if isStartBendA(s) {
			startBendA++
		} else {
			if isStartB(s) {
				startB++
			} else if isEndA(s) {
				endA++
			}
		}
	}

	ans := includeAB
	if startB == 0 && endA == 0 {
		ans += int(math.Max(0, float64(startBendA-1)))
	} else {
		ans += startBendA + int(math.Min(float64(startB), float64(endA)))
	}
	Println(ans)
}
