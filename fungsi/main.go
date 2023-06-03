package main

import (
	"fmt"
)

func main() {
	sekat := getInput()
	potong := getPotong()

	for i := 1; i <= potong; i++ {
		fmt.Printf("Potongan ke- %d\n", i)
		potongSekat(sekat)
		printSekat(sekat)
	}
}

func getInput() []int {
	var n int
	fmt.Scan(&n)

	sekat := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&sekat[i])
	}

	return sekat
}

func getPotong() int {
	var potong int
	fmt.Scan(&potong)
	return potong
}

func potongSekat(sekat []int) {
	for i := 0; i < len(sekat); i++ {
		sekat[i]--
		if sekat[i] < 0 {
			sekat[i] = 0
		}
	}
}

func printSekat(sekat []int) {
	for _, val := range sekat {
		if val < 0 {
			val = 0
		}
		fmt.Println(val)
	}
}
