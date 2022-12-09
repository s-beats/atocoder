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

	n, m, s, t int
)

func main() {
	l := readIntSlice()
	n, m = l[0], l[1]
	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = []int{}
	}
	for i := 0; i < m; i++ {
		l := readIntSlice()
		graph[l[0]] = append(graph[l[0]], l[1])
	}
	l = readIntSlice()
	s, t = l[0], l[1]

	if m == 0 {
		Println(-1)
		return
	}

	// iにj歩目で到達する最短距離
	dist := make([][]int, n+1)
	for i := range dist {
		dist[i] = make([]int, 4)
	}
	dist[s][0] = 1
	que := [][]int{{s, 0}}
	for len(que) > 0 {
		cur := que[0]
		i, j := cur[0], cur[1]
		que = que[1:]
		for _, v := range graph[i] {
			if j == 2 && v == t {
				Println(dist[i][2])
				return
			}
			next := (j % 3) + 1
			if dist[v][next] == 0 {
				dist[v][next] = dist[i][j]
				if j == 3 {
					dist[v][next] += 1
				}
				que = append(que, []int{v, next})
			}
		}
	}
	Println(-1)
}
