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

var layerList, pattyList []int

func solve(n, x int) int {
	if n == 0 {
		return 1
	}
	if x <= n {
		return 0
	}
	if x <= 1+(layerList[n-1]) {
		return solve(n-1, x-1)
	}
	if x == 2+(layerList[n-1]) {
		return pattyList[n-1] + 1
	}
	if x <= 2+(2*layerList[n-1]) {
		return 1 + pattyList[n-1] + solve(n-1, x-2-layerList[n-1])
	}
	return pattyList[n-1] + pattyList[n-1] + 1
}

func main() {
	l := readIntSlice()
	n, x := l[0], l[1]

	layerList = make([]int, n+1)
	layerList[0] = 1
	for i := 1; i < n+1; i++ {
		layerList[i] = (layerList[i-1] * 2) + 3
	}
	pattyList = make([]int, n+1)
	pattyList[0] = 1
	for i := 1; i < n+1; i++ {
		pattyList[i] = (pattyList[i-1] * 2) + 1
	}
	Println(solve(n, x))
}
