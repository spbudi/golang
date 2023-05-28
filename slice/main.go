package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x := scanner.Text()
	evenNames := strings.Split(x, " ")
	result := namaGenap(evenNames)
	fmt.Println(result)

}

func namaGenap(evenNames []string) []string {
	var baru []string
	for _, evenName := range evenNames {
		if len(evenName)%2 == 0 {
			baru = append(baru, evenName)
		}
	}
	return baru
}
