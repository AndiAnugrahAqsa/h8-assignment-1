package main

import (
	"fmt"
	"os"
	"strconv"
)

type Friend struct {
	name    string
	age     int
	address string
	reason  string
}

func main() {
	var argsRaw = os.Args

	var args = argsRaw[1:]

	age, _ := strconv.Atoi(args[1])

	var friend = Friend{
		name:    args[0],
		age:     age,
		address: args[2],
		reason:  args[3],
	}

	fmt.Printf("%+v", friend)
}
