package main

import (
	"bufio"
	"bytes"
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

func readByteSlice() []byte {
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
	return buf
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

type bs = []byte

func main() {
	l := readIntSlice()
	h, w := l[0], l[1]

	s := make([]bs, h)
	t := make([]bs, h)

	for i := 0; i < h; i++ {
		l := readByteSlice()
		s[i] = make(bs, len(l))
		copy(s[i], l)
	}
	for i := 0; i < h; i++ {
		l := readByteSlice()
		t[i] = make(bs, len(l))
		copy(t[i], l)
	}

	// [][]stringを作るより、[]stringのほうが早い？
	// []byteに要素を入れていって最後にstring変換
	c1 := make([]bs, w)
	c2 := make([]bs, w)
	for _, v := range s {
		for j := range v {
			c1[j] = append(c1[j], v[j])
		}
	}
	for _, v := range t {
		for j := range v {
			c2[j] = append(c2[j], v[j])
		}
	}

	// Lessでうまいことやりたい
	sort.Slice(c1, func(i, j int) bool { return string(c1[i]) < string(c1[j]) })
	sort.Slice(c2, func(i, j int) bool { return string(c2[i]) < string(c2[j]) })

	for i, v := range c1 {
		if !bytes.Equal(c2[i], v) {
			Println("No")
			return
		}
	}
	Println("Yes")
}
