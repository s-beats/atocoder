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
		l, p, err := rdr.ReadLine()
		if err != nil {
			panic(err)
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

func toStr(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func isOdd(n int) bool {
	return n&1 == 1
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

var (
	dr = []int{1, -1, 0, 0}
	dc = []int{0, 0, 1, -1}
)

func main() {
	l := readIntSlice()
	h, w := l[0], l[1]
	h--
	w--
	a := make([][]int, h+1)
	que := [][]int{}
	dist := make([][]int, h+1)
	for i := range dist {
		dist[i] = make([]int, w+1)
	}
	for i := 0; i < h+1; i++ {
		buf := make([]int, 0, 16)
		for {
			l, p, err := rdr.ReadLine()
			if err != nil {
				panic(err)
			}
			for ni, v := range l {
				if v == []byte("#")[0] {
					que = append(que, []int{i, ni})
					dist[i][ni] = 1
				}
				buf = append(buf, int(v))
			}
			if !p {
				break
			}
		}
		a[i] = buf
	}

	ans := 0

	for len(que) > 0 {
		cur := que[0]
		que = que[1:]

		// 四方向を走査
		for i := 0; i < 4; i++ {
			r := cur[0] + dr[i]
			c := cur[1] + dc[i]
			if r < 0 || r > h || c < 0 || c > w {
				continue
			}
			if dist[r][c] == 0 {
				dist[r][c] = dist[cur[0]][cur[1]] + 1
				ans = max(ans, dist[r][c])
				que = append(que, []int{r, c})
			}
		}
	}

	if ans == 0 {
		Println(0)
	} else {
		Println(ans - 1)
	}
}
