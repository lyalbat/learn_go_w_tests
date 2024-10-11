package main

import "fmt"

func CustomHello(name string) string {
	const greeting = "hello "
	if(name == ""){
		name = "world"
	}
	return greeting + name
}

func main() {
	fmt.Println(CustomHello("larissa"))
}