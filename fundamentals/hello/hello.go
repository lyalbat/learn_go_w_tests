package main

import "fmt"

const (
	french = "Francais"
	portuguese = "Portugues"
)

func CustomHello(name string) string {
	const greeting = "hello "
	if(name == ""){
		name = "world"
	}
	return greeting + name
}

func InternationHello(language string) (greeting string) {
	switch language {
	case french:
		greeting = "bonjour"
	case portuguese:
		greeting = "bom dia"
	default:
		greeting = "hello"
	}
	return greeting
}

func main() {
	fmt.Println(CustomHello("larissa"))
}