package main

import "fmt"

func main() {
	students := []map[string]string{
		{"name": "Mario", "score": "A"},
		{"name": "Luigi", "score": "B"},
		{"name": "Monster", "score": "C"},
	}

	for _, student := range students{
		fmt.Println(student["name"])
		fmt.Println(student["score"])
	}


}