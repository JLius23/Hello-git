//hola pp
package main

import (
	"fmt"
	"log"
	"keyboard"
)

func main() {
	fmt.Println("Enter a grado:" )
	grade, err := keyboard.GetFloat ()
	if err != nil {
		log.Fatal(err)
	}
var status string 
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("a grado of", grade, "is", status)
}