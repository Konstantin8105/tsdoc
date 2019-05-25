// +build ignore
package main

import (
	"fmt"
	"github.com/Konstantin8105/tsdoc"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot get present folder: %v", err)
		os.Exit(1)
	}

	doc, err := tsdoc.Get(pwd, true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating of documentation from folder `%s`: %v", pwd, err)
		os.Exit(1)
	}

	fmt.Println(doc)
}
