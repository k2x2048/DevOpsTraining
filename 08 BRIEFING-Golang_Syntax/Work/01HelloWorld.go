package main

import ("fmt")

func main() {
	var UserName string
	fmt.Print("write your name: ")
	fmt.Scanln(&UserName)
	if UserName == "" {
		fmt.Println("Hello World !!!\n")
	} else {
		fmt.Println("Hello", UserName, "!!!\n")
	}
}


