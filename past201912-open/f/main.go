package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
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

// 大文字小文字の区別をしないソートができる型
type originalStringSlice []string

func (x originalStringSlice) Len() int { return len(x) }
func (x originalStringSlice) Less(i, j int) bool {
	return strings.ToLower(x[i]) < strings.ToLower(x[j])
}
func (x originalStringSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func (x originalStringSlice) Sort() { sort.Sort(x) }

func main() {
	S := readline()

	SS := []string{}
	startWord := false
	for _, v := range S {
		if unicode.IsUpper(v) {
			if startWord {
				startWord = false
				SS[len(SS)-1] += string(v)
			} else {
				startWord = true
				SS = append(SS, string(v))
			}
		} else {
			SS[len(SS)-1] += string(v)
		}
	}

	sort.Sort(originalStringSlice(SS))

	Println(strings.Join(SS, ""))
}
