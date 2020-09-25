package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func createFile(filename string) {
	if isFileExist(filename) == true {
		e := os.Remove(filename)
		if e != nil {
			log.Fatal(e)
		}

	}
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Gagal membuat file: %v", err)
		return
	}
	defer f.Close()
	min := 1
	max := 30001
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30000; i++ {
		_, err = f.WriteString(fmt.Sprintf("%d\n", rand.Intn(max-min)+min)) // writing...
		if err != nil {
			fmt.Printf("Gagal menulis difile: %v", err)
		}
	}
}

func readFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var s []int

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

func resultSorting(filename string, array []int) {
	if isFileExist(filename) == true {
		e := os.Remove(filename)
		if e != nil {
			log.Fatal(e)
		}
	}

	file, err := os.Create(filename) // creating...
	if err != nil {
		fmt.Printf("Gagal membuat file: %v", err)
		return
	}
	defer file.Close()

	for i := 0; i < len(array); i++ { // Generating...
		_, err = file.WriteString(fmt.Sprintf("%d\n", array[i])) // writing...
		if err != nil {
			fmt.Printf("Gagal menulis difile: %v", err)
		}
	}
}

func isFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	var filename string = "data.txt"
	now := time.Now()
	createFile(filename)
	var data = readFile(filename)
	defer func() {
		fmt.Println("\nEksekusi waktu sorting", time.Now().Sub(now))
		fmt.Printf("Dari jumlah %d data\n\n", len(data))
		fmt.Println("--- Silahkan lihat file hasil.txt ---")
	}()
	fmt.Println("================= ")
	fmt.Println("Shell Sort")
	fmt.Println("================= ")
	shellSort(data)
	resultSorting("hasil.txt", data)
}

func shellSort(array []int) []int {
	var (
		n    = len(array)
		slot = int(math.Floor(float64(n / 2)))
	)
	for slot > 0 {
		for i := slot; i < n; i++ {
			temp := array[i]
			j := i
			for j >= slot && array[j-slot] > temp {
				array[j] = array[j-slot]
				j -= slot
			}
			array[j] = temp
		}
		slot = int(math.Floor(float64(slot / 2)))
	}
	return array
}
