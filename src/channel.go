//  单向通道  
//  数据传递方向是其类型一部分

//  <-chan  接受类型通道  换一句话说可以理解为 只读
//  chan<-  发送类型通道  换一句话说可以理解为 只写

package main

import "fmt"

func ping(pings chan<- string, msg string) {
  pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
  msg := <-pings
  pongs <- msg
}

func main() {
  pings := make(chan string, 1)
  pongs := make(chan string, 1)
  ping(pings, "passed message")
  pong(pings, pongs)
  fmt.Println(<-pongs)
}


//  缓存和非缓冲通道  堵塞问题 或者进入 asleep 状态
//  初始化通道
func main() {
  //  第二个参数 10 表示该通道值在同一时刻最多可以缓冲10个元素值	
  c1 := make(chan int, 10)
	
  //  如果第二个参数被省略了 就表示该通道永远无法缓冲任何元素值 发送给通道的值应该被立刻取走
  //  否则发送方的 goroutine 就会被暂停（或者说阻塞）直到有接收方接受这个值	
  c2 := make(chan string)	
}
  
