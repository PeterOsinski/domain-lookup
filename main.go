package main

import "fmt"
import "sync"
import "time"
import "net"
import "container/heap"

// An IntHeap is a min-heap of ints.
type IntHeap []string

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(string))
}

func (h *IntHeap) Pop() interface{} {

	old := *h
	n := len(old)

	if n == 0 {
		return nil
	}

	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var h = &IntHeap{}

func main() {

	wg := new(sync.WaitGroup)
	var mutex = &sync.Mutex{}

	// Adding routines to workgroup and running then
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(mutex)
	}

	heap.Init(h)

	go getPermutations(4, "qwertyuiopasdfghjklzxcvbnm", mutex)

	wg.Wait()

}

var i int

func worker(mutex *sync.Mutex) {
	for {

		if h.Len() > 0 {

			mutex.Lock()
			addr := heap.Pop(h).(string) + ".pl"
			mutex.Unlock()

			if ip, ok := net.LookupHost(addr); ok == nil {
				i++
				fmt.Printf("%s,%s\n", ip[0], addr)

				if i == 100 {
					fmt.Printf("### %s: Left on heap: %d\n", time.Now().Local(), h.Len())
					i = 0
				}
			}

		}

		time.Sleep(time.Nanosecond)
	}
}
