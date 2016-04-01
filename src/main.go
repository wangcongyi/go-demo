package main

import (
	"fmt"
//"net/http"
)

//func main() {
//	http.HandleFunc("/", handler)
//	http.HandleFunc("/logout", logout)
//	http.ListenAndServe("localhost:9000", nil)
//}
//
//func handler(rw http.ResponseWriter, req *http.Request) {
//	req.ParseForm()
//	if (len(req.Form["name"]) > 0) {
//		fmt.Fprint(rw, "hello ", req.Form["name"][0])
//	}else {
//
//		fmt.Fprint(rw, "Hello world")
//	}
//}
//
//func logout(rw http.ResponseWriter, req *http.Request) {
//	fmt.Fprint(rw, "wahahah")
//}
//var a int = 2
//func main(){
//	b:= 1
//	fmt.Println(b)
//}

//var x, y, z = "haha", 5, 6
//
//func main() {
//	var (
//		a int
//		b bool
//	)
//	fmt.Println(a, b, "\n")
//	fmt.Println(x, y, z)
//}

//const a = 1
//const (
//	text = "123"
//	length = len(text)
//	num = a*20
//)
//
//func main(){
//   fmt.Println(a)
//   fmt.Println(text)
//   fmt.Println(length)
//   fmt.Println(num)
//}

//const (
//	m = iota+1
//	t
//	w
//	f
//	s
//)


/*
  6:   0110
  11:  1011
 ..........
 &    0010
 |    1111
 ^    1101



*/
const (
        B  uint64 =1<<(iota*10)
        KB
        MB
        GB
        TB
)

func main() {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
}












