//   关键字 func 和函数名之间的参数 被称作  接受者
//   将函数与接受者的类型绑定在一起

package main

import "fmt"

type user struct {
	name  string
	age   int
	email string
}

func (u user) notify() {
	fmt.Println(u.name, u.age, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", 20, "bill@email.com"}
	bill.notify()

	// bill := &user{"bill",22,"bill111"}
	// bill.changeEmail("bill122222")
	// bill.notify()
}

//  无语啊  一定要是指针接受者才会改变 。。。。。。。



//   嵌入类型
package main

import "fmt"

type user struct {
	name  string
	email string
}

type admin struct {
	user    // 嵌入类型
	level int
}

func (u *user) notify() {
	fmt.Println(u.name, u.email)
}

func main() {
	ad := admin{
	    user:  user{"king", "king.com"},
	    level: 100,
	}
	ad.notify()           //  内部类型的方法会被提升到外部类型
	ad.user.notify()      //  直接访问内部类型的方法
//  如果外部类型实现了 notify 方法，内部类型的实现就不会被提升。
//  内部类型的值是一直存在的，因此还是可以直接访问内部类型的值来调用没有被提升的内部类型实现方法	
}


//  Mehthos and pointer indirection
package main

import (
	"fmt"
	"math"
)

type V struct {
	x, y float64
}

func (v V) A() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func A(v V) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := V{3, 4}
	fmt.Println(v.A())
	fmt.Println(A(v))

	p := &V{4, 3}
	fmt.Println(p.A())
	fmt.Println(A(*p))
}




