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
//Creating file: data.txt which contains numbers between 1 to 30000 arranged randomly.
func create_file(file_name string) {
	if check_file(file_name) == true {
		
		e := os.Remove(file_name) // if already present, delete the file
		if e != nil {
			log.Fatal(e)
		}

	}
	f, err := os.Create(file_name) // creating ...
	if err != nil {
		fmt.Printf("File not created: %v", err)
		return
	}
	defer f.Close()
	min := 1
	max := 30001
	rand.Seed(time.Now().UnixNano()) 

	for i := 0; i < 30000; i++ {     // Generating numbers...
		_, err = f.WriteString(fmt.Sprintf("%d\n", rand.Intn(max-min)+min)) 
    // writing the numbers in the file
		if err != nil {
			fmt.Printf("Failed to write in the file : %v", err)
		}
	}
}

func read_file(file_name string) []int {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var s []int
	
  //Reading the file data.txt
	scanner := bufio.NewScanner(file) 
	for scanner.Scan() {              //Reading all numbers one by one in a loop
		lineStr := scanner.Text()       
		num, _ := strconv.Atoi(lineStr) //to convert string to int
		s = append(s, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return s
}

//Creating result.txt to store all sorted numbers from data.txt
func result_after_soring(file_name string, array []int) {
	if check_file(file_name) == true {
		
		e := os.Remove(file_name) // Removing already present file, if any
		if e != nil {
			log.Fatal(e)
		}

	}
	f, err := os.Create(file_name) // creating result.txt
	if err != nil {
		fmt.Printf("File not created: %v", err)
		return
	}
	defer f.Close()

	for i := 0; i < len(array); i++ { 
		_, err = f.WriteString(fmt.Sprintf("%d\n", array[i])) // writing sorted numbers in the file - result.txt
		if err != nil {
			fmt.Printf("Failed to write in the file: %v", err)
		}
	}
}

func check_file(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	var file_name string = "data.txt"
	now := time.Now()
	create_file(file_name)
	var data = read_file(file_name)
	defer func() {
		fmt.Println("\nTime Taken for sorting", time.Now().Sub(now))
		fmt.Printf("Data consists of %d numbers\n\n", len(data))
		fmt.Println("--- Result after sorting file 'data.txt' is stored in result.txt ---")
	}()
	fmt.Println("================= ")
	fmt.Println("Insertion Sort")
	fmt.Println("================= ")
	Insertion_sort(data,len(data))
	result_after_soring("result.txt", data)
}

func Insertion_sort(array []int, size int) {
	var j int
    var temp int

    // Iterate over all elements in the array
    for i := 1; i < size; i++ {
        temp = array[i]
        j = i - 1

        /* Pick an element temp and insert it in its correct
        position in sorted array by comparing it with all 
        elements coming before it. Once we find an element
        which is greater than temp, we shift all the greater
        elements to right by one place and temp takes its 
        correct place in the array. */
        for j >= 0 && array[j] > temp {
            array[j + 1] = array[j]
            j--
        }
        
        // Placing temp at its correct position
        array[j + 1] = temp
    }
	
}
