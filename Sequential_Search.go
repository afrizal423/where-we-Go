package main

import (
	"fmt"
	"time"
	"unsafe"
)

func sequential_search(array [5]int, size int, nilai int) bool {
	ketemu := false

	for i := 0; i < size && !ketemu; {
		if array[i] == nilai {
			ketemu = true
		} else {
			i++
		}
	}
	return ketemu

}

func print_array(array [5]int, size int) {
	fmt.Printf("[")

	for i := 0; i < size; i++ {
		fmt.Printf("%d", array[i])

		if i != size-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n")

}

func main() {
	now := time.Now()
	defer func() {
		fmt.Println("Eksekusi waktu ", time.Now().Sub(now))
	}()
	data := [5]int{300, 200, 500, 400, 100}
	size := unsafe.Sizeof(data) / unsafe.Sizeof(data[0])
	// sequential_search
	fmt.Println("================= ")
	fmt.Println("Sequential Search")
	fmt.Println("================= ")
	fmt.Println("Masukkan angka yang dicari: ")
	var cari int
	fmt.Scanln(&cari)
	fmt.Printf("Isi array: \n")
	print_array(data, int(size))
	status_200 := sequential_search(data, int(size), cari)
	fmt.Printf("\nNilai %d %t\n", cari, status_200)
}
