package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type BangunDatar interface{
	Luas() int
}

type Segitiga struct{
	Alas, Tinggi int
}

func (segitiga *Segitiga) Luas() int{
	return segitiga.Alas * segitiga.Tinggi / 2
}

func getLuas(bangundatar BangunDatar){
	fmt.Println("Luas dari Bangun Datar tersebut adalah :",bangundatar.Luas())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	alas, _ := strconv.Atoi(scanner.Text()) 
	scanner.Scan()
	tinggi, _ := strconv.Atoi(scanner.Text()) 
	segitiga := Segitiga{
		Alas: alas,
		Tinggi: tinggi,
	}
	getLuas(&segitiga)

}