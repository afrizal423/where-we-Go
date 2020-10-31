package main

import "fmt"

type element struct {
	npm      string
	name     string
	semester int
	next     *element
  

}

type singleList struct {
	counter int
	head    *element
    tail     *element
}

func initList() *singleList {
	return &singleList{}
}

func (s *singleList) addData() {

	var npm, name string
	var semester int
	fmt.Println("Enter NPM =")
	fmt.Scanln(&npm)
	fmt.Println("Enter name =")
	fmt.Scanln(&name)
	fmt.Println("Enter Semester =")
	fmt.Scanln(&semester)

	element := &element{
		name:     name,
		npm:      npm,
		semester: semester,
	}

	if s.head == nil {
		s.head = element
    s.tail=element
    
	} else {
		element.next = s.head
		s.head = element
    
	}
	s.counter++

	return
}

func (s *singleList) Traverse() error {
	if s.head == nil {
		return fmt.Errorf("LIst is Empty")
	}
	current := s.head
	for current != nil {
		fmt.Println("------- Student Data List -------")
		fmt.Println("NPM : ", current.npm)
		fmt.Println("name : ", current.name)
		fmt.Println("Semester : ", current.semester)
		fmt.Println("-----------------------------------")
		current = current.next
	}
	return nil
}


func (s *singleList) deleteFirstData() error {
	if s.head == nil {
		return fmt.Errorf("LIst is Empty")
	}
  
    if( s.counter == 1){
        s.head=nil
        s.tail=nil
        s.counter--
    } else {
	    s.head = s.head.next
  
	    s.counter--
    }
	return nil

}

func (s *singleList) deleteSingleData() error {
	var npm string
	fmt.Println("Enter NPM that you want to delete =")
	fmt.Scanln(&npm)
    if( s.counter == 1 && s.head.npm==npm){
    s.head=nil
    s.tail=nil
    s.counter--
    return nil
   } 
    temp := s.head
	flag := false
	if(temp.npm==npm){
	    
	    s.head=temp.next
	    s.counter--
	    flag=true
	} else{
	for {
		if temp.next == nil {
			break
		} else if temp.next.npm == npm {

			flag = true
			break
		}
		temp = temp.next
	}
	
	if flag { 
        if(temp.next.next==nil){
            s.tail=temp
            temp.next=nil
    } else{
		    temp.next = temp.next.next
		    s.counter--
        }
	} else{
    
		    fmt.Println("The id you want to delete does not exist")
	    }
	}
  
	return nil
}

func main() {
	singleList := initList()
OUTER:

	fmt.Println("------- MENU - Student Data -------")
	fmt.Println("1. Add data")
	fmt.Println("2. Show data")
	fmt.Println("3. Remove From Front")
	fmt.Println("4. Delete Single Data")
	fmt.Println("--------------------------------------")
	fmt.Println("Total Data : ", singleList.counter)
	fmt.Println("--------------------------------------")
	fmt.Println("\nInput:")
	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		singleList.addData()
		goto OUTER
	case 2:
		singleList.Traverse()
		goto OUTER
	case 3:
		singleList.deleteFirstData()
		fmt.Println("-------FRONT ELEMENT DELETED--------\n")
		goto OUTER
	case 4:
		singleList.deleteSingleData()
        fmt.Println("-------ELEMENT DELETED--------\n")
		goto OUTER
	}

}