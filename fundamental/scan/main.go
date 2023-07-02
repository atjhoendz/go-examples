package main

import "fmt"

func main() {
	/*
		scan
		- only store space-separated value
	*/
	fmt.Println("example scan standard input")

	fmt.Print("input your firstname: ")
	var firstname, fullname string
	// scan separated space input
	fmt.Scan(&firstname)

	fmt.Print("input your fullname: ")
	// scan input stop by newline
	fmt.Scanln(&fullname)

	fmt.Printf("your firstname is %s\n", firstname)
	fmt.Printf("your fullname is %s\n", fullname)
}
