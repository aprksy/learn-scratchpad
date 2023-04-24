package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func makeSelection() {
	fmt.Printf("select scenario (%d):\n", len(scenarios))
	for i, item := range scenarios {
		fmt.Printf("%d. %s\n", i+1, item.Name)
	}
	fmt.Print("enter your selection: ")
	input := bufio.NewScanner(os.Stdin)
	var selectionStr string
	if input.Scan() {
		selectionStr = input.Text()
		if value, err := strconv.Atoi(selectionStr); err != nil {
			fmt.Println("please enter number from above")
			makeSelection()
		} else {
			fmt.Printf("your selection is number '%d. %s'\n", value, scenarios[value-1].Name)
			if err := scenarios[value-1].Action(); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func main() {
	makeSelection()
}
