package main

import "fmt"

func main() {
	var count int
	fmt.Scanln(&count)

	for i := count; i >= 1; i--{
		fmt.Printf("%d I will become go developer \n", i)
	}
}