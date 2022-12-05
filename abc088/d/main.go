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

func readRuneSlice() []rune {
	line := readline()
	slice := make([]rune, len(line))
	for i, v := range line {
		slice[i] = v
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

func bfs(grid [][]rune) [][]int {
	que := make([][]int, 0, 5)
	que = append(que, []int{0, 0})

	dist := make([][]int, len(grid))
	for i := range dist {
		dist[i] = make([]int, len(grid[0]))
	}
	dist[0][0] = 1

	updateOrNoop := func(curVal, row, column int) {
		if grid[row][column] == rune('#') {
			return
		}
		if dist[row][column] != 0 {
			return
		}
		dist[row][column] += curVal + 1
		que = append(que, []int{row, column})
	}

	for len(que) != 0 {
		for _, v := range que {
			curVal := dist[v[0]][v[1]]

			// x + 1
			if v[1] != len(grid[0])-1 {
				updateOrNoop(curVal, v[0], v[1]+1)
			}
			// x - 1
			if v[1] != 0 {
				updateOrNoop(curVal, v[0], v[1]-1)
			}

			// y + 1
			if v[0] != len(grid)-1 {
				updateOrNoop(curVal, v[0]+1, v[1])
			}
			if v[0] != 0 {
				updateOrNoop(curVal, v[0]-1, v[1])
			}

			que = que[1:]
		}
	}
	return dist
}

func main() {
	l := readIntSlice()
	h, w := l[0], l[1]

	whiteCount := 0
	grid := make([][]rune, h)
	for i := 0; i < h; i++ {
		line := readline()
		slice := make([]rune, len(line))
		for i, v := range line {
			if v == rune('.') {
				whiteCount++
			}
			slice[i] = v
		}
		grid[i] = slice
	}

	goal := bfs(grid)[h-1][w-1]
	if goal == 0 {
		Println(-1)
	} else {
		Println(whiteCount - goal)
	}
}
