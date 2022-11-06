package main

import "math"

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
