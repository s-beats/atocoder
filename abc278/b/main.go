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
	l := readIntSlice()
	h, m := l[0], l[1]

	for {
		hs := strconv.Itoa(h)
		ms := strconv.Itoa(m)

		if len(hs) == 1 {
			hs = "0" + hs
		}
		if len(ms) == 1 {
			ms = "0" + ms
		}

		splitH := strings.Split(hs, "")
		splitM := strings.Split(ms, "")
		newH := splitH[0] + splitM[0]
		newM := splitH[1] + splitM[1]

		newHInt := toInt(newH)
		newMInt := toInt(newM)
		if newHInt < 24 && newMInt < 60 {
			Println(strconv.Itoa(h) + " " + strconv.Itoa(m))
			return
		}

		if m == 59 {
			h++
			if h == 24 {
				h = 0
			}
			m = 0
		} else {
			m++
		}
	}
}
