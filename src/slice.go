package main

import "fmt"

func main() {

	/*
	     GO 数组 一旦声明就无法改变长度和数据类型

	     长度与数据类型相同才能赋值
	*/

	var arr1 [10]int                   // [0 0 0 0 0 0 0 0 0 0]

	arr2 := [10]int{}                  // [0 0 0 0 0 0 0 0 0 0]

	arr3 := [5]int{1, 2, 3, 4, 5}      // [1 2 3 4 5]

	arr4 := [...]int{}                 // []

	arr5 := [...]int{3:50}             // [0 0 0 50]

	arr6 := [5]int{1:20, 2:20}         //  [0 20 20 0 0]


	/*

	     GO 切片  把底层数组切出一部分
         相对于数组  使用切片可以按需增加切片的容量
         切片访问的元素个数          长度  len()
         切片允许增长到的元素个数     容量  cap()
	*/


	slice1 := make([]string, 5)        // [     ]

	slice2 := make([]string, 3, 5)     // [   ]

	slice3 := []int{1, 2, 3, 4, 5}     // [1 2 3 4 5]

	slice4 := []string{99:"slice"}     // [    ...     slice]

	source := []int{1, 2, 3, 4, 5}     // 使用切片创建切片  两个切片共享同一个底层数组 如果一个切片修改了，另一个切片也会影响
	slice5 := source[1:3]              // [2 3]         -->> 包括起始索引 不包括结束索引
	slice6 := source[1:]               // [2 3 4 5]     -->> 取后面i个值
	slice7 := source[:3]               // [1 2 3]       -->> 取前面i个值



	s := []int{1,2,3,4,5}
	n := s[1:3]

	fmt.Println(n)      // [2 3]


	n = append(n,22)
	fmt.Println(n)      //  [2 3 22]
	fmt.Println(s)      //  [1 2 3 22 5]

	n = append(n,223)
	fmt.Println(n)     //  [2 3 22 223]
	fmt.Println(s)     //  [1 2 3 22 223]

	fmt.Println(cap(n))    // 4
	fmt.Println(len(n))    // 4

	n = append(n,2234)
	fmt.Println(n)    //  [2 3 22 223 2234]
	fmt.Println(s)    //  [1 2 3 22 223]


	fmt.Println(cap(n))     // 8
	fmt.Println(len(n))     // 5

	n = append(n,22345)

	fmt.Println(cap(n))     // 8
	fmt.Println(len(n))     // 6


	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)
	fmt.Println(slice7)

}
