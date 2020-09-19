package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func buat_file(nama_file string) {
	if cek_file(nama_file) == true {
		// hapus file
		e := os.Remove(nama_file) // hapus file
		if e != nil {
			log.Fatal(e)
		}

	}
	f, err := os.Create(nama_file) // creating...
	if err != nil {
		fmt.Printf("Gagal membuat file: %v", err)
		return
	}
	defer f.Close()
	min := 1
	max := 30001
	rand.Seed(time.Now().UnixNano()) // ambil satuan waktu untuk di random int
	for i := 0; i < 30000; i++ {     // Generating...
		_, err = f.WriteString(fmt.Sprintf("%d\n", rand.Intn(max-min)+min)) // writing...
		if err != nil {
			fmt.Printf("Gagal menulis difile: %v", err)
		}
	}
}

func baca_file(nama_file string) []int {
	file, err := os.Open(nama_file)
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

func hasil_sort(nama_file string, array []int) {
	if cek_file(nama_file) == true {
		// hapus file
		e := os.Remove(nama_file) // hapus file
		if e != nil {
			log.Fatal(e)
		}

	}
	f, err := os.Create(nama_file) // creating...
	if err != nil {
		fmt.Printf("Gagal membuat file: %v", err)
		return
	}
	defer f.Close()

	for i := 0; i < len(array); i++ { // Generating...
		_, err = f.WriteString(fmt.Sprintf("%d\n", array[i])) // writing...
		if err != nil {
			fmt.Printf("Gagal menulis difile: %v", err)
		}
	}
}

func cek_file(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	var nama_file string = "data.txt"
	now := time.Now()
	buat_file(nama_file)
	var data = baca_file(nama_file)
	defer func() {
		fmt.Println("\nEksekusi waktu sorting", time.Now().Sub(now))
		fmt.Printf("Dari jumlah %d data\n\n", len(data))
		fmt.Println("--- Silahkan lihat file hasil.txt ---")
	}()
	fmt.Println("================= ")
	fmt.Println("Bubble Sort")
	fmt.Println("================= ")
	bubblesort(data)
	hasil_sort("hasil.txt", data)
}

func bubblesort(array []int) {
	var (
		n      = len(array)
		sorted = false
	)
	for !sorted {
		pindahkan := false
		for i := 0; i < n-1; i++ {
			if array[i] > array[i+1] {
				array[i+1], array[i] = array[i], array[i+1]
				pindahkan = true
			}
		}
		if !pindahkan {
			sorted = true
		}
		n = n - 1
	}
}
