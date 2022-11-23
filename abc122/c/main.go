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

// 累積和
func main() {
	l := readIntSlice()
	_, Q := l[0], l[1]
	S := strings.Split(readline(), "")

	cs := make([]int, len(S)+1)
	for i := range S {
		if i > 0 && S[i-1] == "A" && S[i] == "C" {
			cs[i+1] = cs[i] + 1
		} else {
			cs[i+1] = cs[i]
		}
	}

	for i := 0; i < Q; i++ {
		l := readIntSlice()
		Println(cs[l[1]] - cs[l[0]])
	}
}
