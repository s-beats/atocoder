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
	_, q := l[0], l[1]

	m := make(map[int]map[int]bool)

	for i := 0; i < q; i++ {
		l := readIntSlice()
		t, ai, bi := l[0], l[1], l[2]
		if t == 1 {
			if _, ok := m[ai]; !ok {
				m[ai] = make(map[int]bool)
			}
			m[ai][bi] = true
		} else if t == 2 {
			if _, ok := m[ai]; !ok {
				m[ai] = make(map[int]bool)
			}
			m[ai][bi] = false
		} else {
			_, aok := m[ai]
			_, bok := m[bi]
			if aok && bok && m[ai][bi] && m[bi][ai] {
				Println("Yes")
			} else {
				Println("No")
			}
		}
	}
}
