package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)
	wg.Add(2)
	go player("A", court)
	go player("B", court)
	court <- 1
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court

		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("Player %s Missed n is %d\n", name, n)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d n is %d\n", name, ball, n)
		ball++
		court <- ball
	}
}
