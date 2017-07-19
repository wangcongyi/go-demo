//  由于go语言没有没有class的概念
//  结构struct承担了面向对象的功能

type vertex struct {
	x, y int
}

var (
	v1 = vertex{1, 2}
	v2 = vertex{x:1}
	v3 = vertex{}
	v4 = &vertex{1, 2}
)

func main() {
  fmt.Println(v1)                  //  {1,2}
	fmt.Println(v2)                  //  {1,0}
	fmt.Println(v3)                  //  {0,0}
	fmt.Println(v4)                  //  &{1,2}    特殊的前缀 & 返回一个指向结构体的指针。
}



*********************************************
*********************************************


type human struct {
  Sex  int
}

type teacher struct {
  human
  Name  string
  Age   int
}

type person struct {
  Name string
  Age  int
  Contenxt struct {
    Phone,City string
  }
}

func main () {
  a := teacher{Name:"king", Age:20, human:human{Sex:0}}
  b := person{Name:"King",Age:30}
  b.Contenxt.Phone = "2232323"
  b.Contenxt.City = "SZ"
  fmt.Println(a)
  fmt.Println(b)
}

//////////
// 含有相同‘key'的 struct
type A struct {
  B
  Name string
}

type B struct {
  Name string
}

func main () {
  a := A{Name:"A",B:B{Name:"B"}}

  fmt.Println(a.name,a.B.name)

}


