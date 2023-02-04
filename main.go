package main

import "fmt"

type Student struct {
	firstName, lastName string
	id                  int
	age                 int
	school              string
	gpa                 float64
}

type Teacher struct {
	firstName, lastName string
	age                 int
	subject             string
}

func main() {
	if studentFull().gpa >= 2.7 && teacherFull().subject == "ReserchMethods" {
		fmt.Println(studentFull().firstName, studentFull().lastName, "Good GPA!!")
		fmt.Print("Teacher - ", teacherFull())
	} else {
		fmt.Print(studentFull(), " - GPA is not good!!(")
	}
}

func studentFull() Student {
	Stud1 := Student{firstName: "Eldar", lastName: "Sharapiyev", id: 200103253, age: 19, school: "SDU", gpa: 3}
	return Stud1
}

func teacherFull() Teacher {
	Teacher1 := Teacher{firstName: "Gulnura", lastName: "", age: 21, subject: "ReserchMethods"}
	return Teacher1
}
