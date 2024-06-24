package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)
	studentGrades["Rohit"] = 95
	studentGrades["Rahul"] = 85
	studentGrades["Raman"] = 88
	fmt.Println("Marks of Raman:", studentGrades["Raman"])

	//To change a value
	// studentGrades["Rahul"] = 87
	// fmt.Println("Marks of Rahul",studentGrades["Rahul"])
	
	//To delete
	delete(studentGrades, "Rahul")
	fmt.Println("Marks of Rahul:", studentGrades["Rahul"])

	//To check if a key is present
	grades,exists := studentGrades["Rohit"]
	fmt.Println(grades,exists)
}
