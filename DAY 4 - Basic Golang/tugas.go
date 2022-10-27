/*
Nama 		: Leonardo Dwigantoro
Program 	: The “Rect n Tri” Angle
Deskripsi	: Buatlah sebuah fungsi dengan sebuah parameter angka. Jika angkanya ganjil maka program akan loop untuk print angka tersebut membentuk Triangle. Jika angka genap maka program akan loop membentuk Rectangle. Bonus kalau angkanya diinput di console
Aqua Developer Academy Batch 2 Bandung
*/

package main

import "fmt"

func RectNTri(input int) {

	// inisiasi variabel untuk looping
	var i, j int

	// jika input merupakan bilangan genap
	if input%2 == 0 {
		fmt.Println("=====================================")
		fmt.Printf("%d adalah bilangan genap\n", input)
		fmt.Println("=====================================")

		// looping untuk menampilkan persegi dengan sisi merupakan besaran inputan
		for i = 1; i <= input; i++ {
			for j = 1; j <= input; j++ {
				fmt.Printf("*")
			}
			fmt.Printf("\n")
		}

		// jika input merupakan bilangan ganjil
	} else {
		fmt.Println("=====================================")
		fmt.Printf("%d adalah bilangan ganjil\n", input)
		fmt.Println("=====================================")

		// looping untuk menampilkan segitiga dengan tinggi merupakan besaran inputan dan alas merupakan besaran inputan dikurangi 1
		for i = 1; i <= input; i++ {
			for j = 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
		fmt.Println("=====================================")
	}

}

func main() {

	// iniaiasi variabel untuk menampung inputan
	var inputAngka int

	// menampilkan text untuk input angka
	fmt.Println("=====================================")
	fmt.Printf("Masukan input berupa angka: ")

	// mengambil input dari user
	fmt.Scanln(&inputAngka)

	// memanggil fungsi RectNTri
	RectNTri(inputAngka)

}
