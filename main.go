package main

import "fmt"
import "sync"
import "time"
import "net"
import "flag"
import "runtime"

var domains = []string{"pl","com","eu","com.pl"}

func worker(workerId int) {
	for {
		addr := getPermutation()
		for _, domain := range domains {
			domain = addr + "." + domain
			if ip, ok := net.ResolveIPAddr("ip4", domain); ok == nil {
				fmt.Printf("%s\n", domain)
			}
		}
	}
	time.Sleep(time.Nanosecond)
}

type CF struct {
	workerNum 		int
	addressLength 	int
}

var cf = new(CF)

func cli() {
	flag.IntVar(&cf.workerNum, "worker_num", 1000, "Number of workers")
	flag.IntVar(&cf.addressLength, "address_len", 4, "Max number of characters in generated domains")
	
	flag.Parse()	
}

func main() {
	
	runtime.GOMAXPROCS(3)
	
	cli()
	
	wg := new(sync.WaitGroup)
	initPermutations(cf.addressLength, "qwertyuiopasdfghjklzxcvbnm")

	// Adding routines to workgroup and running then
	for i := 0; i < cf.workerNum; i++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()

}
