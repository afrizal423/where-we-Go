package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func sequential_search(array []int, size int, nilai int) bool {
	/*
		Kita inisialisasikan variable ketemu bertipe boolean dengan nilai false
	*/
	ketemu := false
	/*
		lalu kita looping semua data yang ada pada array data, sesuai jumlah arraynya.
		Dengan kondisi jika i kurang dari size dan variable ketemu bernilai false
	*/
	for i := 0; i < size && !ketemu; {
		/*
			Kita cek nilai yang ada pada array tersebut apakah sama dengan nilai yang dicari
		*/
		if array[i] == nilai {
			/*
				Jika nilainya yang ada pada array tersebut sama dengan nilai yang dicari,
				maka kita akan menset variable ketemu menjadi true
			*/
			ketemu = true
		} else {
			/*
				Jika nilainya yang ada pada array tersebut tidak sama dengan nilai yang dicari,
				maka akan menambahkan nilai +1
			*/
			i++
		}
	}
	return ketemu

}

func print_array(array []int, size int) {
	fmt.Printf("[")

	for i := 0; i < size; i++ {
		fmt.Printf("%d", array[i])

		if i != size-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n")

}

func baca_file() []int {
	file, err := os.Open("h.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var s []int
	// printSlice(s)
	scanner := bufio.NewScanner(file) // inisialisasi membaca file
	for scanner.Scan() {              // loop reading datanya
		lineStr := scanner.Text()       // ambil 1 data
		num, _ := strconv.Atoi(lineStr) // ubah ke int
		s = append(s, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return s
}

func main() {
	now := time.Now()
	var data = baca_file()
	defer func() {
		fmt.Println("\nEksekusi waktu ", time.Now().Sub(now))
		fmt.Printf("Dari jumlah %d data\n\n", len(data))
	}()
	fmt.Println("================= ")
	fmt.Println("Sequential Search")
	fmt.Println("================= ")
	fmt.Println("Masukkan angka yang dicari: ")
	// kita beri variable cari sebagai penampung inputan
	var cari int
	// lalu kita akan memproses inputan
	fmt.Scanln(&cari)
	// panggil fungsinya
	// fmt.Printf("Isi array: \n")
	// print_array(data, len(data)) // tampilkan data jika mau
	status_200 := sequential_search(data, len(data), cari)
	fmt.Printf("\nNilai %d %t\n", cari, status_200)
}
