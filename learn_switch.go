package main

import "fmt"

//switch语句
func main()  {
  switch a := 10; a {
  case 10:
    fmt.Println("golang switch so bad")
  case 20:
    fmt.Println("hahaha...")
  }
  // 等价于如下写法
  // a := 10
  // switch  {
  // case 10:
  //
  // case:20
  //
  // }
}
