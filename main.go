package main

import (
	"fmt"
	"os"
)

const GH_ACTIONS_INPUT_PREF = "INPUT_"

func main() {
	fmt.Println("Hello, World!")

	input1 := os.Getenv(fmt.Sprintf("%s%s", GH_ACTIONS_INPUT_PREF, "INPUT1"))
	input2 := os.Getenv(fmt.Sprintf("%s%s", GH_ACTIONS_INPUT_PREF, "INPUT2"))

	fmt.Printf("Input1: %s\n", input1)
	fmt.Printf("Input2: %s\n", input2)
}
