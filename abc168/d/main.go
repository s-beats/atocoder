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
	n, m := l[0], l[1]

	grid := make([][]int, n+1)
	for i := 0; i < m; i++ {
		l = readIntSlice()
		grid[l[0]] = append(grid[l[0]], l[1])
		grid[l[1]] = append(grid[l[1]], l[0])
	}

	ans := make([]int, n+1)
	queue := []int{}
	queue = append(queue, 1)
	visit := make([]int, n+1)
	visit[1] = 1

	for len(queue) != 0 {
		// pop
		elm := queue[0]
		queue = queue[1:]
		for _, v := range grid[elm] {
			if visit[v] != 0 {
				continue
			}

			visit[v] = 1
			ans[v] = elm
			queue = append(queue, v)
		}
	}

	buf := []byte{}
	buf = append(buf, []byte("Yes\n")...)
	for _, v := range ans[2:] {
		if v == 0 {
			Println("No")
			return
		}
		buf = append(buf, []byte(strconv.Itoa(v)+"\n")...)
	}
	Println(string(buf[:len(buf)-1]))
}
