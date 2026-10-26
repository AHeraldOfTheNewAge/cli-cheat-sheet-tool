package main

import (
	"fmt"
	"os"
)

//TODO -> Make an error function!!!

func readCheatSheet(fileName string) {
	info, err := os.Stat("cheatsheets")

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Error: cheetsheet tool broken! Cheat-sheets folder does not exist.")

			return
		}

		fmt.Println("Error: cheetsheet tool broken!", err)
	}

	if !info.IsDir() {
		fmt.Println("Error: cheetsheet tool broken! The cheat sheets should be stored in a folder, how is this not a folder!?")
	}

	data, err := os.ReadFile("cheatsheets/" + fileName + ".txt")

	if err != nil {
		fmt.Println("Error: Cannot read the cheatsheet!", err)

		return
	}

	fmt.Println("\n" + string(data))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide at least one argument! \n For example: cheetsheet docker!")

		return
	}

	var argVal string = os.Args[1]

	readCheatSheet(argVal)
}
