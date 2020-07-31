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


//  缓存和非缓冲通道  堵塞问题 或者进入 asleep 状态  另写
  
