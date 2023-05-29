package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func main() {
	month :=[12]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei" ,
		"juni" ,
		"juli" ,
		"agustus" ,
		"september" ,
		"november" ,
		"desember" ,
	}

	sales := map[string]int{
		"januari": 2836,
		"februari": 3282,
		"maret": 787,
		"april": 6238,
		"mei": 1992,
		"juni": 824,
		"juli": 2903,
		"agustus": 602,
		"september": 930,
		"oktoner": 1002,
		"november": 922,
		"desember": 3219,
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	bulan1,_ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	bulan2,_ := strconv.Atoi(scanner.Text())

	for i := bulan1 -1; i < bulan2; i++{
		fmt.Printf("Bulan %s : %d tusuk\n", month[i], sales[month[i]])
	}
}