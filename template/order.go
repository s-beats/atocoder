package main

import (
	"fmt"
	"strconv"
	"time"
)

// n * 10 でTLEしない計算量

func joinString() string {
	var n int = 1e4
	now := time.Now()
	s := ""
	for i := 0; i < n; i++ {
		s += strconv.Itoa(i)
	}
	fmt.Println(time.Since(now).String())
	return s
}

func appendBytes() string {
	var n int = 1e7
	now := time.Now()
	s := []byte{}
	for i := 0; i < n; i++ {
		s = append(s, []byte(strconv.Itoa(i))...)
	}
	fmt.Println(time.Since(now).String())
	return string(s)
}

func compareBytes() bool {
	var n int = 1e8
	now := time.Now()
	s1 := make([]byte, n)
	s2 := make([]byte, n)
	ans := false
	for i := 0; i < n; i++ {
		ans = s1[i] == s2[i]
	}
	fmt.Println(time.Since(now).String())
	return ans
}

func appendInt() {
	var n int = 1e7
	now := time.Now()
	s := []int{}
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	fmt.Println(time.Since(now).String())
}
