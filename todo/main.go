package main

import (
	"fmt"
	"os"
)

type Task struct {
	ID   int
	Name string
}

var tasks []Task
var nextID int

func main() {
	fmt.Println("Welcome to the ToDo List Application")
	for {
		fmt.Println("\n1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Delete All Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask()
		case 2:
			listTasks()
		case 3:
			deleteTask()
		case 4:
			deleteAllTask()
		case 5:
			fmt.Println("Thank you for using the ToDo List Application")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again")

		}
	}
}

func addTask(){
	var name string
	fmt.Print("Enter task name: ")
	fmt.Scanln(&name)
	nextID++
	task := Task{ID: nextID, Name: name}
	tasks = append(tasks, task)
	fmt.Println("Task Added Successfully")
}

func listTasks(){
	fmt.Println("List Of Tasks")
	for _, task := range tasks {
		fmt.Printf("%d: %s\n", task.ID, task.Name)
	}
}

func deleteTask(){
	fmt.Print("Enter task ID to delete: ")
	var id int;
	fmt.Scanln(&id)
	
	for i , task := range tasks {
		if task.ID == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			
			fmt.Println("Task Deleted Successfully")
			return
		}
	}

}

func deleteAllTask(){
	tasks = nil;
	nextID = 0;
	fmt.Println("All Tasks Deleted Successfully")
}

