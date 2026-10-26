package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var assetsDir = "cheetsheets"

func isCheatSheetsFolder() bool {
	info, err := os.Stat(assetsDir)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Error: cheetsheet tool broken! Cheat-sheets folder does not exist.")

			return false
		}

		fmt.Println("Error: cheetsheet tool broken!", err)

		return false
	}

	if !info.IsDir() {
		fmt.Println("Error: cheetsheet tool broken! The cheat sheets should be stored in a folder, how is this not a folder!?")

		return false
	}

	return true
}

func printAvailable() {
	if !isCheatSheetsFolder() {
		return
	}

	entries, err := os.ReadDir(assetsDir)

	if err != nil {
		fmt.Println("Error: cheetsheet tool broken! Cannot read in the cheat sheets folder", err)
	}

	fmt.Println("Available cheat sheets: \n")

	for _, entry := range entries {
		if entry.IsDir() { // Why would a folder be here?!.. anyway..
			continue
		}

		fmt.Println("*", strings.ToLower(strings.Replace(entry.Name(), ".txt", "", 1)))
	}

	fmt.Println("")
}

func printCheatSheet(cheatSheetName string) {
	if !isCheatSheetsFolder() {
		return
	}

	data, err := os.ReadFile(filepath.Join(assetsDir, cheatSheetName+".txt"))

	if err != nil {
		fmt.Println("Error: Cannot read the cheatsheet!", err)

		return
	}

	const (
		ColorReset = "\033[0m"
		ColorCyan  = "\033[36m"
	)

	fmt.Printf("%s==== %s ====%s\n", ColorCyan, cheatSheetName+" cheat-sheet", ColorReset)
	fmt.Println("\n" + string(data))
	fmt.Printf("\n%s================%s\n", ColorCyan, ColorReset)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide at least one argument! \n For example: cheetsheet help")

		return
	}

	var argVal string = os.Args[1]

	argVal = strings.ToLower(argVal)

	if argVal == "help" {
		fmt.Println("This small cli tool prints a little cheat-sheet straight to your cli." +
			"\nTo list avalialbe cheat sheets type cheetsheet available, and it will list the available cheat sheets!" +
			"\nTo list a cheat sheet just type cheetsheet <name> and it will print the desired cheat sheet in the terminal!\n")

		return
	}

	if argVal == "available" {
		printAvailable()

		return
	}

	printCheatSheet(argVal)
}
