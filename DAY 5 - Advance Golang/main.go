/*
Nama 		: Leonardo Dwigantoro
Program 	: Tugas Week 1 Soal 1
Aqua Developer Academy Batch 2 Bandung
*/
package main

import (
	"fmt"
	"sort"
)

// membuat struct product
type Product struct {
	Name  string
	Price int
}

func main() {

	var products = []Product{
		0: {
			Name:  "Benih Lele",
			Price: 50000,
		},
		1: {
			Name:  "Pakan lele cap menara",
			Price: 25000,
		},
		2: {
			Name:  "Probiotik A",
			Price: 75000,
		},
		3: {
			Name:  "Probiotik Nila B",
			Price: 10000,
		},
		4: {
			Name:  "Pakan Nila",
			Price: 20000,
		},
		5: {
			Name:  "Benih Nila",
			Price: 20000,
		},
		6: {
			Name:  "Cupang",
			Price: 5000,
		},
		7: {
			Name:  "Benih Nila",
			Price: 30000,
		},
		8: {
			Name:  "Benih Cupang",
			Price: 10000,
		},
		9: {
			Name:  "Probiotik B",
			Price: 10000,
		},
	}

	// memanggil fungsi MaximumItemsCanBeBoughtOnce
	MaximumItemsCanBeBoughtOnce(products)

	// mencari produk termahal dan termurah
	FindMaxAndMin(products)

	// mencari produk dengan harga 10000
	FindSpecificProduct(products)

}

// fungsi untuk mencari produk terbanyak yang bisa dibeli dengan uang 100000 dimana masing-masing produk hanya bisa dibeli 1 kali
func MaximumItemsCanBeBoughtOnce(products []Product) {
	fmt.Println("====================================")
	fmt.Println("Total produk dengan harga dibawah  Rp 100.000 :")
	fmt.Println("Harga total 100000")

	// sorting nilai dari terkecil ke terbesar
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})

	var total int
	for _, product := range products {
		// jika nilai sudah lebih besar dari 100000 maka akan dihentikan
		if total+product.Price > 100000 {
			break
		}
		total += product.Price
		fmt.Println(product.Name, "-", product.Price)
	}
}

// fungsi untuk mencari produk termahal dan termurah
func FindMaxAndMin(products []Product) {
	max := Product{}
	min := products[0]
	fmt.Println("====================================")
	for _, product := range products {
		if product.Price > max.Price {
			max = product
		}
		if product.Price < min.Price {
			min = product
		}
	}
	fmt.Printf("Daftar produk termurah: %s Rp %d \n", min.Name, min.Price)
	fmt.Printf("Daftar produk termahal: %s Rp %d \n", max.Name, max.Price)
}

// fungsi untuk mencari produk-produk dengan harga 10000
func FindSpecificProduct(products []Product) {
	value := Product{}
	fmt.Println("====================================")
	for _, product := range products {
		if product.Price == 10000 {
			value = product
			fmt.Printf("%s - %d \n", value.Name, value.Price)
		}
	}
	fmt.Println("====================================")
}
