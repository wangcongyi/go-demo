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



// 嵌入类型
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
}

