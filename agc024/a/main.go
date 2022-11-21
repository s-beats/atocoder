package main

import (
	"bufio"
	. "fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// 1行の文字列がデフォルトを超えることがあるのでサイズを明示的に指定
var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

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

func pow(x, y int) int {
	bigx := big.NewInt(int64(x))
	bigy := big.NewInt(int64(y))
	return int(bigx.Exp(bigx, bigy, nil).Int64())
}

func main() {
	l := readIntSlice()
	a, b, _, k := l[0], l[1], l[2], l[3]
	unfair := math.Pow10(18)
	ans := a - b
	if k > 0 {
		ans *= int(pow(-1, k))
	}
	ansAbs := math.Abs(float64(ans))
	if ansAbs > unfair {
		Println("Unfair")
	} else {
		Println(ans)
	}
}
