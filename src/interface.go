//////////////////////////////////////////////
// an empty interface may hold values of any type.
// emppt interface are used by code that handles values of unknown type.
package main

import "fmt"

func d(i interface{}) {
	fmt.Printf("%v %T\n", i, i)
}

func main() {
	var i interface{}
	d(i)

	i = 24323
	d(i)

	i = "hello"
	d(i)
}

///////////////////////////////////////////////////////
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

func (t *T) M() {
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}

func d(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	i = &T{"hello"}
	d(i)
	i.M()

	i = F(math.Pi)
	d(i)
	i.M()
}
