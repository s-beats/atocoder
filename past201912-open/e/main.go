package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
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
	L1 := readIntSlice()
	N, Q := L1[0], L1[1]

	// キーのユーザーがフォローしているユーザーのリストを表す
	follow := make(map[int][]bool, 0)
	for i := 0; i <= N; i++ {
		follow[i] = make([]bool, N+1)
	}

	for i := 0; i < Q; i++ {
		l := readIntSlice()

		if l[0] == 1 {
			solve(follow, l[0], l[1], l[2])
		} else {
			solve(follow, l[0], l[1], 0)
		}
	}

	for i := 1; i <= N; i++ {
		var l string
		for j := 1; j <= N; j++ {
			if follow[i][j] {
				l += "Y"
			} else {
				l += "N"
			}
		}
		Println(l)
	}
}

func solve(follow map[int][]bool, n, a, b int) {
	switch n {
	case 1:
		follow[a][b] = true
	case 2:
		f := follow
		for k, v := range f {
			if k == a {
				continue
			}

			if v[a] {
				follow[a][k] = true
			}
		}
	case 3:
		fa := make([]bool, len(follow[a]))
		copy(fa, follow[a])
		for i, v := range fa {
			if v {
				for ii, vv := range follow[i] {
					if ii == a {
						continue
					}
					if vv {
						follow[a][ii] = true
					}
				}
			}
		}
	}

}
