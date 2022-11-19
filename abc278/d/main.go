package main

import (
	"bufio"
	. "fmt"
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
	_ = toInt(readline())
	a := readIntSlice()
	q := toInt(readline())

	var latestOne *int
	var latestSlice map[int]int
	for i := 0; i < q; i++ {
		l := readIntSlice()
		t, x := l[0], l[1]
		if t == 1 {
			latestOne = new(int)
			*latestOne = x
			latestSlice = make(map[int]int)
		} else if t == 2 {
			y := l[2]
			if latestOne != nil {
				if latestSlice[x-1] > 0 {
					latestSlice[x-1] += y
				} else {
					latestSlice[x-1] += *latestOne + y
				}
			} else {
				a[x-1] += y
			}
		} else {
			if latestOne != nil {
				if latestSlice[x-1] == 0 {
					latestSlice[x-1] += *latestOne
				}
				Println(latestSlice[x-1])
			} else {
				Println(a[x-1])
			}
		}
	}
}
