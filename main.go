package main

import (
	"os"

	"github.com/berkeleycole/gopractice/lessons"
)

func main() {
	switch os.Args[1] {
	case "1":
		fallthrough
	case "HelloWorld":
		lessons.HelloWorld()
	case "2":
		lessons.ValuesAndVariables()
	case "3":
		lessons.Constants()
	}
}
