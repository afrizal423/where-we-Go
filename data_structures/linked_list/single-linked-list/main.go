package main

import "fmt"

type element struct {
	npm      string
	nama     string
	semester int
	next     *element
}

type singleList struct {
	counter int
	head    *element
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

	if s.head == nil {
		s.head = element
	} else {
		element.next = s.head
		s.head = element
	}
	s.counter++

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
	}
	return nil
}

func (s *singleList) HapusDiDepan() error {
	if s.head == nil {
		return fmt.Errorf("Listnya kosong")
	}
	s.head = s.head.next
	s.counter--
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
			s.counter--
			pPre.next = current.next
			return nil
		}

		pPre = current
		current = current.next
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
	fmt.Println("Total data yang ada : ", singleList.counter)
	fmt.Println("--------------------------------------")
	fmt.Println("\nmasukan pilihan:")
	var pilih int
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
