package main

import (
	"sync"
	"runtime"
	"fmt"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)

	fmt.Println("Start")
	go s("A")
	go s("B")

	fmt.Println("Waiting")
	wg.Wait()

	fmt.Println("All Over")

}

func s(p string) {
	defer wg.Done()

next:
	for o := 2; o < 5000; o++ {
		for i := 2; i < o; i++ {
			if o%i == 0 {
				continue next
			}
		}
		fmt.Printf("%s: %d\n", p, o)
	}
	fmt.Println("Over", p)
}
