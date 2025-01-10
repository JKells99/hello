package main

import (
	"fmt"
)

var students []Student
var courses []Course
var nextStudentId int
var nextCourseId int

func main() {
	fmt.Println("Welcome To Studnet System!")
	for {
		fmt.Println("\n1. Add Student")
		fmt.Println("2. List Students")
		fmt.Println("3. Delete Student")
		fmt.Print("Enter Choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addStudent()

		case 2:
			printStudents()

		case 3:
			deleteStudent()

		default:
			fmt.Println("Invalid Choice. Please Try Again")
		}
	}
}

func addStudent() {
	var name, grade string
	fmt.Println("Enter Student Name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter Student Grade: ")
	fmt.Scanln(&grade)
	nextStudentId++
	student := Student{ID: nextStudentId, Name: name, Grade: grade}
	students = append(students, student)
	fmt.Println("Student Added Successfully!")
	fmt.Println(students)
}

func printStudents() {
	fmt.Println("List Of Students")
	for _, student := range students {
		fmt.Println("Student Details ", student.Name)
		fmt.Println("ID: " ,student.ID)
		fmt.Println("Name: " ,student.Name)
		fmt.Println("Grade: " ,student.Grade)
		fmt.Println()

	}
}

func deleteStudent(){
	fmt.Println("enter name to delete")
	var name string
	fmt.Scanln(&name)
	for i, student := range students{
		if student.Name == name{
			// This works by creating a new slice that contains all the elements before the element we want to remove and all the elements after the element we want to remove.
			students = append(students[:i], students[i+1:]...)
			
			fmt.Println("Student Deleted Successfully")
		}
}
}

