//  GO语言中的 map 相对于 其他语言中的 hashes 或者 dicts


package main

improt "fmt"

func main(){

   // 使用映射字面量创建，初始长度根据初始化时指定的键值对来确定
   // 映射的键 可以是任何值 
   dict := map[string]string{"red":"#da1337", "orange":"#e95a22"}
   
   ******************************************
   // 创建一个空映射，通过指定合适的类型赋值
   colors := map[string]string{}
   colors["red"] = "#da1337"
   
   ******************************************
   //声明一个未初始化的 nil 映射,但 nil 映射不能用于存储键值对
   var colors map[string]string
   colors["red"] = "#da1137"      Runtime Error: panic: runtime error: assigment to entry in nil map
   
   
  
   ***********************************************
   //从映射获取值并判断是否存在 
   colors := map[string]string{"blue":"aa", "red":"bb"}
   value, exists := colors["blue"]
   if exists {
     fmt.Println(value)     // aa
     fmt.Println(exists)    // true
   }
   
   
   ************************************************
   // 使用 range 迭代映射 
   colors := map[string]int{
   // 相同的 key 值会报错 duplicate key "xxx" in map literal
      "a":1,
      "b":2,
      "c":3,
   }
    // 从映射中删除一项
     delete(colors, "a")
   
   for key, value := range colors {
   //  直接打印出来 顺序并非与初始化的一致
       fmt.Printf(key, value)     
   }
   
   
   ************************************************
   arr := [5]int{1, 2, 3, 4, 5}

	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(arr)
	fmt.Println(slice)
   
   
}
