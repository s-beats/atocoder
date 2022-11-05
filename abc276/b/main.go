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

const BUFSIZE = 1000000

var rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)

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

func mod(n1, n2 int) int {
	res := (n1 + n2) % n2
	if res < 0 {
		res += n2
	}
	return res
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func popBack(l *[]interface{}) interface{} {
	e := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return e
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
}

func isOdd(n int) bool {
	return n&1 == 1
}

func isEven(n int) bool {
	return n&1 == 0
}

type mem map[int]int

func main() {
	l := readIntSlice()
	N, M := l[0], l[1]

	ans := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		ans[i] = []int{}
	}

	for i := 0; i < M; i++ {
		l = readIntSlice()
		ans[l[0]] = append(ans[l[0]], l[1])
		ans[l[1]] = append(ans[l[1]], l[0])
	}

	for i, v := range ans {
		if i == 0 {
			continue
		}
		sort.IntSlice(v).Sort()
	}

	for i, v := range ans {
		if i == 0 {
			continue
		}
		Print(len(v))
		if len(v) == 0 {
			Println()
			continue
		}
		for _, vv := range v {
			Print(" ", vv)
		}
		Println()
	}
}
