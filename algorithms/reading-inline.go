package main

import (
	"fmt"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)

	// Second reader
	// fmt.Print("Enter text for a NewScanner: ")
	var name string
	// var unit string
	// var amount string
	// var temp string

	// fmt.Scanln(&name, &temp, &amount, &unit)
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	if name != "" {
		fmt.Printf("Welcome %v to goland \n", name)
	}

	fmt.Print("Please enter 2 number to sum: ")
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	fmt.Println(sum(a, b, c))

	//	fmt.Print("Enter text: ")
	//	text := ""
	//	fmt.Scanln(&text)
	//	fmt.Println(text)

	//	fmt.Print("Enter text: ")
	//	text := ""
	//	fmt.Scanln(&text)
	//	fmt.Println(text)

	// for scanner.Scan() {
	// 	text += scanner.Text()
	// 	fmt.Print(textt)
	// }

	// if scanner.Err() != nil {
	// 	// handle error
	// }

	// fmt.Print("Enter text2: ")
	// text2 := ""
	// fmt.Scanln(text2)
	// fmt.Println(text2)
	//
	// ln := ""
	// fmt.Sscanln("%v", ln)
	// fmt.Println(ln)
}

func sum(a, b, c int) int {
	return a + b + c
}
