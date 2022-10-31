package main

import (
	"bufio"
	"fmt"
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
			fmt.Println(e.Error())
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

func visitedTarget(start, target int, elements []int, skipStartNegative bool) bool {
	dp := make(map[int]bool)

	dp[start] = true
	if skipStartNegative {
		dp[-start] = true
	}

	for _, v := range elements {
		// ひとつ前のステップのみを保持
		tmp := make(map[int]bool)
		for k := range dp {
			tmp[k+v] = true
			tmp[k-v] = true
		}
		dp = tmp
	}

	return dp[target]
}

func main() {
	NS := readIntSlice()
	_, x, y := NS[0], NS[1], NS[2]

	A := readIntSlice()

	odds := make([]int, 0, len(A)/2+1)
	evens := make([]int, 0, len(A)/2+1)

	// Aiが奇数ならx, 偶数ならy
	for i, v := range A {
		if isOdd(i + 1) {
			odds = append(odds, v)
		} else {
			evens = append(evens, v)
		}
	}

	// X軸の場合は、startの負の場合をスキップ
	// 問題の条件により、必ず正の方向に進むため
	if visitedTarget(odds[0], x, odds[1:], false) && visitedTarget(0, y, evens, true) {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}
