package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukan milik kelereng si A :")
	scanner.Scan()
	giver, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Masukan milik kelereng si B :")
	scanner.Scan()
	receiver, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Masukan jumlah kelereng yang akan diberkan :")
	scanner.Scan()
	marble, _ := strconv.Atoi(scanner.Text())

	giveMarbles(&giver, &receiver, marble)

	if giver < marble {
		fmt.Println("Kelereng si A tidak cukup untuk dibagikan")
	} else {
		fmt.Printf("Sisa Kelereng si A adalah : %d \n", giver)
		fmt.Printf("Total kelereng si B adalah : %d \n", receiver)
	}

}

func giveMarbles(giver *int, receiver *int, marble int) {
	if *giver < marble {
		fmt.Println("Kelereng tidak cukup untuk dibagikan")
		return
	} else if *giver == marble {
		*giver -= marble
		*receiver += marble
		fmt.Println("Kelereng sudah habis dibagikan")
		return
	} else {
		*giver -= marble
		*receiver += marble
	}
}
