package main

import (
	"GoStudyProject/interfacesingo/interfaceonconsumersite/consumer"
	"fmt"
)

func main() {

	a := consumer.Article{Title: "How to learn french in 15 minutes"}
	fmt.Printf("The title is %s\n", a)

}
