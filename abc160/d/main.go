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

func main() {
	l := readIntSlice()
	n, x, y := l[0], l[1], l[2]
	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = []int{}
	}
	for i := 1; i <= n; i++ {
		if i != n {
			graph[i] = append(graph[i], i+1)
			graph[i+1] = append(graph[i+1], i)
		}
		if i == x {
			graph[i] = append(graph[i], y)
		} else if i == y {
			graph[i] = append(graph[i], x)
		}
	}

	ans := make([]int, n+1)
	type pair struct{ a, b int }
	solvedShortest := make(map[pair]bool)

	for i := 1; i <= n; i++ {
		dist := make([]int, n+1)
		dist[i] = 0
		visit := make([]int, n+1)
		visit[i] = 1
		que := []int{}
		que = append(que, i)
		for len(que) != 0 {
			cu := que[0]
			que = que[1:]
			for _, v := range graph[cu] {
				if visit[v] == 0 {
					dist[v] = dist[cu] + 1
					visit[v] = 1
					que = append(que, v)

					p := pair{
						a: min(i, v),
						b: max(i, v),
					}
					if !solvedShortest[p] {
						ans[dist[cu]+1]++
						solvedShortest[p] = true
					}
				}
			}
		}
	}

	for _, v := range ans[1 : len(ans)-1] {
		Println(v)
	}
}
