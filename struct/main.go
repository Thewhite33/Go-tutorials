package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	RollNo    int
}

type Contact struct {
	Phone int
	Email string
}

type Class struct {
	Student_Details Student
	Student_Contact Contact
}

func main() {
	var student Student
	student.FirstName = "Raman"
	student.LastName = "Singh"
	student.RollNo = 32223
	//fmt.Println(student)
}
