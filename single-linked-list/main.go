package main

import "fmt"

type element struct {
	npm      string
	nama     string
	semester int
	next     *element
}

type singleList struct {
	jumlah_data int
	head        *element
}

func initList() *singleList {
	return &singleList{}
}

func (s *singleList) tmbhData() {
	print("\033[H\033[2J")
	var npm, nama string
	var semester int
	fmt.Println("Masukkan NPM =")
	fmt.Scanln(&npm)
	fmt.Println("Masukkan Nama =")
	fmt.Scanln(&nama)
	fmt.Println("Masukkan Semester =")
	fmt.Scanln(&semester)

	element := &element{
		nama:     nama,
		npm:      npm,
		semester: semester,
	}

	// fmt.Println(element)

	if s.head == nil {
		s.head = element
	} else {
		element.next = s.head
		s.head = element
	}
	s.jumlah_data++

	return
}

func (s *singleList) Traverse() error {
	if s.head == nil {
		return fmt.Errorf("Listnya kosong")
	}
	current := s.head
	for current != nil {
		fmt.Println("------- List Data Mahasiswa -------")
		fmt.Println("NPM : ", current.npm)
		fmt.Println("Nama : ", current.nama)
		fmt.Println("Semester : ", current.semester)
		fmt.Println("-----------------------------------")
		current = current.next
		// fmt.Printf("%+v\n", current)
	}
	return nil
}

func (s *singleList) HapusDiDepan() error {
	if s.head == nil {
		return fmt.Errorf("Listnya kosong")
	}
	s.head = s.head.next
	s.jumlah_data--
	return nil
}

func (s *singleList) HapusSingleData() error {
	var npm string
	fmt.Println("Masukkan NPM yang ingin dihapus =")
	fmt.Scanln(&npm)

	if s.head == nil {
		return fmt.Errorf("Listnya kosong")
	}
	current := s.head
	var pPre *element
	for current != nil {
		if current.npm == npm {
			fmt.Println("------- List Data Mahasiswa -------")
			fmt.Println("NPM : ", current.npm)
			fmt.Println("Nama : ", current.nama)
			fmt.Println("Semester : ", current.semester)
			fmt.Println("-----------------------------------")
			// s.head = current.next
			s.jumlah_data--
			// pPre := current
			// anu := current.next
			// fmt.Println(anu, pPre)
			// pPre.next = anu.next
			// pPre = nil
			pPre.next = current.next
			return nil
		}

		// pPre := current
		pPre = current
		current = current.next
		// fmt.Println(current)
	}

	return nil
}

func main() {
	singleList := initList()
	print("\033[H\033[2J")
OUTER:

	fmt.Println("------- Aplikasi Data Mahasiswa -------")
	fmt.Println("1. Tambah data")
	fmt.Println("2. Tampilkan data")
	fmt.Println("3. Hapus Data Didepan")
	fmt.Println("4. Hapus Single Data")
	fmt.Println("--------------------------------------")
	fmt.Println("Total data yang ada : ", singleList.jumlah_data)
	fmt.Println("--------------------------------------")
	fmt.Println("\nmasukan pilihan:")
	var pilih int
	// lalu kita akan memproses inputan
	fmt.Scanln(&pilih)

	switch pilih {
	case 1:
		singleList.tmbhData()
		goto OUTER
	case 2:
		singleList.Traverse()
		goto OUTER
	case 3:
		singleList.HapusDiDepan()
		goto OUTER
	case 4:
		singleList.HapusSingleData()
		goto OUTER
	}

}
