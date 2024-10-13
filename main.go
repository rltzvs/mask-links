package main

import (
	"fmt"
	"go_project/linkmasker"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file> [output_file]")
		return
	}

	inputFile := os.Args[1]
	outputFile := "output.txt"
	if len(os.Args) >= 3 {
		outputFile = os.Args[2]
	}

	prod := linkmasker.NewFileProducer(inputFile)
	pres := linkmasker.NewFilePresenter(outputFile)
	svc := linkmasker.NewService(prod, pres)

	if err := svc.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
