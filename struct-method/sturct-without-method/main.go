package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Penduduk struct {
	nama string
	umur int
	alamat string
}

func main() {
	var penduduk1 Penduduk
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	penduduk1.nama = scanner.Text()
	scanner.Scan()
	penduduk1.umur, _ = strconv.Atoi(scanner.Text()) 
	scanner.Scan()
	penduduk1.alamat = scanner.Text()

	fmt.Printf("Hello, my name is %s. Im %d years old. I live in %s", penduduk1.nama, penduduk1.umur, penduduk1.alamat)

}