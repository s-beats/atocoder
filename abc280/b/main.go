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

func main() {
	n := toInt(readline())
	s := readIntSlice()

	ans := make([]int, n)
	ans[0] = s[0]
	sum := ans[0]
	for i := 1; i < n; i++ {
		ans[i] = s[i] - sum
		sum += ans[i]
	}

	ansStr := ""
	for _, v := range ans {
		ansStr += strconv.Itoa(v) + " "
	}
	Println(strings.TrimRight(ansStr, " "))
}
