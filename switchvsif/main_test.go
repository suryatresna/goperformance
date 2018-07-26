package main

import (
	"fmt"
	"testing"
)

var words = []string{"super", "califragi", "listic", "expiali", "docius"}
var wlen = len(words)

var numbers = []int{1, 2, 3, 4, 5}
var ilen = len(numbers)

func BenchmarkSwitchString(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		switch words[i%wlen] {
		case "super":
			m++
		case "califragi":
			m++
		case "listic":
			m++
		case "expiali":
			m++
		case "docius":
			m++
		}
	}
	fmt.Println(m)
}

func BenchmarkIfString(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		w := words[i%wlen]
		if w == "super" {
			m++
		} else if w == "califragi" {
			m++
		} else if w == "listic" {
			m++
		} else if w == "expiali" {
			m++
		} else if w == "docius" {
			m++
		}
	}
	fmt.Println(m)
}

func BenchmarkSwitchInt(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		switch numbers[i%ilen] {
		case 1:
			m++
		case 2:
			m++
		case 3:
			m++
		case 4:
			m++
		case 5:
			m++
		}
	}
	fmt.Println(m)
}

func BenchmarkIfInt(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		n := numbers[i%ilen]
		if n == 1 {
			m++
		} else if n == 2 {
			m++
		} else if n == 3 {
			m++
		} else if n == 4 {
			m++
		} else if n == 5 {
			m++
		}
	}
	fmt.Println(m)
}
