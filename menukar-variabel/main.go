// package main

// import "fmt"

// func main() {
// 	var a, b string
// 	fmt.Scanln(&a)
// 	fmt.Scanln(&b)

// 	// solusi 1
// 	// temp := a
// 	// a = b
// 	// b = temp

// 	// solusi 2
// 	a, b = b, a
// 	fmt.Println(a, b)
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	kata1, _ := reader.ReadString('\n')
	kata1 = strings.TrimSpace(kata1)

	kata2, _ := reader.ReadString('\n')
	kata2 = strings.TrimSpace(kata2)

	// solusi 1
	// tmp := kata1
	// kata1 = kata2
	// kata2 = tmp

	// solusi 2
	// kata1, kata2 = kata2, kata1

	// solusi 3
	kata1, kata2 = swap(kata1, kata2)

	fmt.Println(kata1, kata2)
}

func swap(parameter1, parameter2 string) (string, string) {
	return parameter2, parameter1
}
