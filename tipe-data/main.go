package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x int
	var y float64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ = strconv.ParseFloat(scanner.Text(), 64)

	result1 := x + int(y)
	result2 := float64(x) - y
	result3 := x * int(y)

	fmt.Printf("%d \n", result1)
	fmt.Printf("%.2f \n", result2)
	fmt.Printf("%.d \n", result3)

}
