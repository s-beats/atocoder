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

func main() {
	S := readline()

	Println(solve(S))
}

func solve(s string) string {
	words := []string{}
	for i := 0; i < len(s); {
		j := i + 1
		for j < len(s) && unicode.IsLower(rune(s[j])) {
			j++
		}

		wordRune := []rune(s[i : j+1])
		wordRune[0] = unicode.ToLower(wordRune[0])
		wordRune[len(wordRune)-1] = unicode.ToLower(wordRune[len(wordRune)-1])
		words = append(words, string(wordRune))

		i = j + 1
	}

	sort.Strings(words)

	for i := 0; i < len(words); i++ {
		wordRune := []rune(words[i])
		wordRune[0] = unicode.ToUpper(wordRune[0])
		wordRune[len(wordRune)-1] = unicode.ToUpper(wordRune[len(wordRune)-1])
		words[i] = string(wordRune)
	}

	return strings.Join(words, "")
}
