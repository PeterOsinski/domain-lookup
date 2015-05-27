package main

import "math"
import "strings"

var positions []int
var chars []string
var current float64

func inc(pos int) bool {

	if pos < 0 {
		return false
	}

	if positions[pos] == len(chars)-1 {
		inc(pos - 1)
		positions[pos] = 0
	} else {
		positions[pos]++
	}

	return true
}

var a float64
var b float64

func initPermutations(capacity int, entropia string) {
	positions = make([]int, capacity)
	chars = strings.Split(entropia, "")

	a = float64(capacity)
	b = float64(len(entropia))
}

func getPermutation() string{

	if current >= math.Pow(b, a) {
		return ""
	}

	var str = make([]string, len(positions))
	for k, v := range positions {
		if v > 0 {
			str[k] = chars[v-1]
		}
	}
	inc(len(positions) - 1)
	current++
	
	return strings.Join(str, "")
}
