package main

import (
	"fmt"

	"github.com/Abraxas-365/gpto/pkg/utils"
)

func main() {
	// Initialize the Golang parser

	_, err := utils.NewFunctionIndexer(".")
	if err != nil {
		fmt.Println("Error while walking the directory:", err)
	}
}
