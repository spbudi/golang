package main

import "fmt"

func main() {
	// var count int
	// fmt.Scanln(&count)

	// for i := count; i >= 1; i--{
	// 	fmt.Printf("%d I will become go developer \n", i)
	// }

	// for i := 1; i <= 5; i++ {
	// 	for j := 1; j <= 6-i; j++{
	// 		fmt.Printf("%d ",j)
	// 	}
	// 	fmt.Println()
	// }

	// var counter int = 0

	// for i := 0; i < 5; i++{
	// 	for j := 5; j > i; j--{
	// 		fmt.Printf("%d ", j)
	// 	}
	// 	fmt.Println()
	// }

	// var i int = 0

	// for i < 5{
	// 	fmt.Println("Angka", i)
	// 	i++
	// }
	outerLoop:

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop // kalau pakai outerLoop i akan berhenti jika bernilai 3, kalau break saja akan lanjut ke iterasi berikutnya
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}