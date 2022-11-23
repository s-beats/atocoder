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

func log(n int) int {
	return int(math.Log(float64(n)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	_ = toInt(readline())
	lineElms := strings.Split(readline(), " ")
	a := make([]int, 0)
	containEven := false
	for _, v := range lineElms {
		containEven = containEven || !isOdd(toInt(v))
		a = append(a, toInt(v))
	}
	if !containEven {
		Println(0)
		return
	}

	ans := 0
	for _, v := range a {
		if isOdd(v) {
			continue
		}

		for v%2 == 0 {
			v /= 2
			ans++
		}
	}
	Println(ans)
	return
}
