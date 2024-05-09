package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error no file")
		os.Exit(1)
	}

	_, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	panic("implement me!")
}
