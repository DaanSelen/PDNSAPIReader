package main

import (
	"fmt"
	"os"
	"strings"
)

var ()

func main() {
	printMainStartScreen()
	printMainCommandScreen()
}

func printMainStartScreen() {
	printLines()
	fmt.Println("Welcome to the PDNSAPIReader.")
	fmt.Println()
	fmt.Println("Press any key to continue...")
	printLines()
	fmt.Scanln()
}

func printMainCommandScreen() {
	printLines()
	fmt.Println("Available commands:")
	fmt.Println()
	fmt.Print("SearchDomain (Retrieves a list of all domains that fit the searchkey)\nShowDomain (Retrieves all Records for a given domain)\nsetTTL (Set Time To Live)\n\nExit (Immediatly exit the program)")
	fmt.Println()
	printLines()
	fmt.Print("Enter the command you want: ")
	var prefCommand string
	fmt.Scanln(&prefCommand)
	selectInput(strings.ToLower(prefCommand))
}

func selectInput(trialCommand string) {
	fmt.Println()
	switch trialCommand {
	case "searchdomain":
		printLines()
	case "showdomain":
		printLines()
	case "setttl":
		printLines()
	case "exit":
		os.Exit(0)
	default:
		printMainCommandScreen()
	}
}

func searchDomain() {

}

func showDomain() {
}

func printLines() {
	fmt.Println(strings.Join([]string{strings.Repeat("=", 80)}, " "))
}
