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
	nama := scanner.Text()
	nama = strings.TrimSpace(nama)
	scanner.Scan()
	alamat := scanner.Text()
	alamat = strings.TrimSpace(alamat)
	scanner.Scan()
	email := scanner.Text()
	email = strings.TrimSpace(alamat)

	fmt.Printf("Nama saya adalah %s, alamat rumah saya adalah %s, dan email saya adalah %s", nama, alamat, email)
}
