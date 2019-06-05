package main

import (
	"fmt"
	"sort"
)

func ArraySource(a ...int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, v := range a {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

func S(in <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		var a []int
		for v := range in {
			a = append(a, v)
		}
		sort.Ints(a)
		for _, v := range a {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func Merge(int1, int2 <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		v1, c1 := <-int1
		v2, c2 := <-int2

		for c1 || c2 {
			if !c2 || (c1 && v1 <= v2) {
				ch <- v1
				v1, c1 = <-int1
			} else {
				ch <- v2
				v2, c2 = <-int2
			}
		}
		close(ch)
	}()

	return ch
}

func main() {
	p := Merge(S(ArraySource(7, 9, 2, 4, 5, 8, 0, 3, 6)), S(ArraySource(32, 34, 56, 65, 7, 4, 65, 67, 455, 76, 45, 6567, 345, )))

	for v := range p {
		fmt.Println(v)
	}
}
