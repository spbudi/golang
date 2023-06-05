package main

import (
	"fmt"
)

type Student struct {
	name string
	age int
}
func (s *Student) Introduction() {
	fmt.Printf("Hello, my name is %s. Im %d years old\n", s.name, s.age)
}

func main(){
	student1 := Student{name: "Sammy", age: 17}
	student1.Introduction()
}

