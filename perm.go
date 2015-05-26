package main

import "math"
import "strings"
import "container/heap"
import "sync"

var positions []int
var chars []string

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

func getPermutations(capacity int, entropia string, mutex *sync.Mutex) {

	positions = make([]int, capacity)
	chars = strings.Split(entropia, "")

	a := float64(capacity)
	b := float64(len(entropia))
	var i float64

	for i = 0; i < math.Pow(b, a); {

		var str = make([]string, capacity)
		for k, v := range positions {
			str[k] = chars[v]
		}

		mutex.Lock()
		heap.Push(h, strings.Join(str, ""))
		mutex.Unlock()

		inc(len(positions) - 1)
		i++
	}
}
