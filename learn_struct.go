package main

import "fmt"

type Student struct {
  name string
  sex int
}

func main()  {
  student := Student{"xiaoming",1}
  fmt.Println(student.name)
  fmt.Println(student.sex)
}
