/////// 并不会修改原来的值  传递值的拷贝
func main(){
  a,b := 1,2
  A(a,b)
  fmt.Println(a,b)   ////    1 2
}

func A(s ...int){
   s[0] = 22
   s[1] = 33 
   fmt.Println(s)    ////  [22 33]
}

/////// 把整个切片传递过去（传递的是内存地址） 会影响道切片本身
func main() {
	s := []int{1, 2, 3, 4, 5}
	A(s)
	fmt.Println(s)
}

func A(s []int) {
	s[0] = 232321

	fmt.Println(s)  ///////  [232321 2 3 4 5]
}

/////本质上都是拷贝传递  但是拷贝的对象不一样（内存地址的copy 和 值的copy）

/////////////
////////
// 下面的demo就是传递内存地址 。从而可以改变原来的值
func main(){
  a := 11
  A(&a)
  fmt.Println(a)
}

func A(a *int){
  *a = 232
  fmt.Println(*a)
}


//GO语言版 闭包

func main() {
	f := c(23)
	fmt.Println(f(2323))
	fmt.Println(f(233323))
}

func c(x int) (func(y int) int) {   
    return func(y int) int {
	    fmt.Println(&x)   //  print the same value twice
		return x + y
	}
}

//////    来一个简单一点的 闭包 好了  

func intSeq() func() int{
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())    // 1
	fmt.Println(nextInt())    // 2
	fmt.Println(nextInt())    // 3
}

