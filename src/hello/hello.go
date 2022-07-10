package main

import "fmt"

func fullname(fname, lname string) string {
	fmt.Println("Calculating....")
	return fname + " " + lname
}

func main() {
	var fname string
	var lname string
	fmt.Println("Enter first name: ")
	fmt.Scanln(&fname)
	fmt.Println("Enter last name: ")
	fmt.Scanln(&lname)

	fmt.Println("Your name is: ", fullname(fname, lname))
}
