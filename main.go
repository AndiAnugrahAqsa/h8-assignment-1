package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	number  int
	name    string
	age     int
	address string
	reason  string
}

func printStudent(students []Student, number int) {
	for _, student := range students {
		if student.number == number {
			fmt.Printf("%+v", student)
			return
		}
	}
	fmt.Printf("%+v", "student not found")
}

func main() {

	var students = []Student{
		{
			number:  1,
			name:    "Rian",
			age:     20,
			address: "Jakarta",
			reason:  "Ingin fokus belajar Golang",
		},
		{
			number:  2,
			name:    "Fikri",
			age:     22,
			address: "Bandung",
			reason:  "Memperdalam bahasa Golang",
		},
		{
			number:  3,
			name:    "Ilham",
			age:     21,
			address: "Surabaya",
			reason:  "Menambah stack di bidang backend",
		},
	}

	var argsRaw = os.Args

	var args = argsRaw[1:]

	if len(args) == 0 {
		fmt.Println("argument is needed")
		return
	}

	number, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("argument must be integer")
		return
	}

	printStudent(students, number)
}
