package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/optclblast/bytecode-vm/instructions"
	"github.com/optclblast/bytecode-vm/stack"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error no file")
		os.Exit(1)
	}

	srcFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	stack := stack.Allocstack(1024)

	scanner := bufio.NewScanner(srcFile)

	for scanner.Scan() {
		row := (strings.Split(scanner.Text(), ";")[0])
		operation := strings.Split(row, " ")[0]
		args := strings.Split(row, " ")[1:]

		//fmt.Printf("Operation: %s | Args: %v\n", operation, args)

		instructions.InstructionsSet[operation].Fn(stack, args)
	}
}
