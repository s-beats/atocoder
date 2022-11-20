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

func main() {
	unfair := math.Pow10(18)
	l := readIntSlice()
	a, b, _, k := l[0], l[1], l[2], l[3]
	ans := a - b
	if k&1 == 1 {
		ans = b - a
	} else {
		ans = a - b
	}
	ansAbs := math.Abs(float64(ans))
	if ansAbs > unfair {
		Println("Unfair")
	} else {
		Println(ans)
	}
}
