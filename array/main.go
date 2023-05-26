package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// start READ kapasitas
	scanner.Scan()
	kapasitas, _ := strconv.Atoi(scanner.Text())
	// end READ kapasitas

	// start INITIALIZE
	lariks := make([]int, kapasitas)
	// end INITIALIZE

	// start READ elemen
	scanner.Scan()
	elemen := scanner.Text()
	pisahElemens := strings.Split(elemen, " ")
	// end READ elemen
	
	for i, pisahElemen := range pisahElemens {
		x, _ := strconv.Atoi(string(pisahElemen))
		lariks[i] = x
	}

	for _, larik := range lariks {
		if larik%2 == 0 && larik != 0 {
			fmt.Println("Angka genapnya adalah:",larik)
		}
	}
}
