package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	selectedIndex = 0
	urlArray      = [2]string{"https://powerdns.systemec.nl/rest/rest.php", "https://teslatest.systemec.nl/rest/rest.php"}
	selectedURL   = urlArray[selectedIndex]
)

func main() {
	printMainStartScreen()
	printMainCommandScreen()
}

func printMainStartScreen() {
	printLines()
	fmt.Println()
	fmt.Println("Welcome to the PDNSAPIReader.")
	fmt.Println()
	printLines()
	pressAny()
}

func printMainCommandScreen() {

	printLines()
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println()
	fmt.Println("COMMANDS:\nSearchDomain (Retrieves a list of all domains that fit the searchkey)\nShowDomain (Retrieves all Records for a given domain)\nsetTTL (Set Time To Live)\n\nExit (Immediatly exit the program)")
	fmt.Println()
	fmt.Println("CLIENT SETTINGS MANAGEMENT:\nURLChange (Option to change between experimental and production API Servers)\nURLStatus (Checks which environment it's going to call (Default: Production))")
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
		searchDomain()
	case "showdomain":
		showDomain()
	case "setttl":
		setTTL()
	case "exit":
		os.Exit(0)
	case "urlchange", "uc":
		urlChange()
	case "urlstatus", "us":
		urlStatus()
	}
	printMainCommandScreen()
}

func searchDomain() {
	printLines()
	fmt.Println()
	fmt.Print("What is your searchkey? ")
	var searchKeyInput string
	fmt.Scanln(&searchKeyInput)
	searchKey := confirm(searchKeyInput)
	fmt.Println(searchKey)
	fmt.Println()
	printLines()
}

func showDomain() {
	printLines()
	fmt.Println()

	fmt.Println()
	printLines()
}

func setTTL() {
	printLines()
	fmt.Println()

	fmt.Println()
	printLines()
}

func confirm(trialInput string) string {
	printLines()
	fmt.Println()
	fmt.Print("Your input is:", trialInput)
	fmt.Println()
	printLines()
	fmt.Print("is this correct? (Y/n) ")
	var confirmInput string
	fmt.Scanln(&confirmInput)
	if strings.ToLower(confirmInput) == "n" {
		fmt.Print("Type in the correct input: ")
		var correctedInput string
		fmt.Scanln(&correctedInput)
		return correctedInput
	} else {
		return trialInput
	}
}

func urlChange() {
	if selectedIndex == 0 {
		selectedURL = urlArray[1]
		selectedIndex = 1
	} else {
		selectedURL = urlArray[0]
		selectedIndex = 0
	}
	printLines()
	fmt.Println()
	fmt.Println("URL CHANGE SUCCES. Current:", selectedURL)
	fmt.Println()
	printLines()
	pressAny()
}

func urlStatus() {
	printLines()
	fmt.Println()
	fmt.Println("Current URL:", selectedURL)
	fmt.Println()
	printLines()
	pressAny()
}

func printLines() {
	fmt.Println(strings.Join([]string{strings.Repeat("=", 80)}, " "))
}

func pressAny() {
	fmt.Println("Press any key to continue...")
	fmt.Scanln()
}
