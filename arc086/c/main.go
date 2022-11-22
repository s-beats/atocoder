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
	N := toInt(readline())
	abs := func(n int) int {
		return int(math.Abs(float64(n)))
	}
	cur := []int{0, 0}
	prev := 0
	for i := 0; i < N; i++ {
		l := readIntSlice()
		t, x, y := l[0], l[1], l[2]
		xabs := abs(x - cur[0])
		yabs := abs(y - cur[1])
		elapse := t - prev
		if elapse < xabs+yabs {
			Println("No")
			return
		}
		if elapse&1 != (xabs+yabs)&1 {
			Println("No")
			return
		}
		prev = t
		cur[0] = x
		cur[1] = y
	}
	Println("Yes")
}
