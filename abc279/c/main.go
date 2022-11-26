package main

import (
	"bufio"
	. "fmt"
	"math"
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

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func isOdd(n int) bool {
	return n&1 == 1
}

func main() {
	l := readIntSlice()
	h, w := l[0], l[1]

	s := make([]string, h)
	t := make([]string, h)

	for i := 0; i < h; i++ {
		s[i] = readline()
	}
	for i := 0; i < h; i++ {
		t[i] = readline()
	}

	// [][]stringを作るより、[]stringのほうが早い？
	// []byteに要素を入れていって最後にstring変換
	c1 := make([]string, w)
	c2 := make([]string, w)
	for i := range c1 {
		bs := make([]byte, h)
		for j, v := range s {
			bs[j] = v[i]
		}
		c1[i] = string(bs)
	}
	for i := range c2 {
		bs := make([]byte, h)
		for j, v := range t {
			bs[j] = v[i]
		}
		c2[i] = string(bs)
	}

	sort.Strings(c1)
	sort.Strings(c2)

	for i, v := range c1 {
		if v != c2[i] {
			Println("No")
			return
		}
	}
	Println("Yes")
}
