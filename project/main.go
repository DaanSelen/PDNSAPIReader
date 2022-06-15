package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	selectedIndex = 0
	urlArray      = [2]string{"https://powerdns.systemec.nl/rest/rest.php", "https://teslatest.systemec.nl/rest/rest.php"}
	selectedURL   = urlArray[selectedIndex]
	user          string
	password      string
)

func main() {
	printMainStartScreen()
	printMainCommandScreen()
}

func printMainStartScreen() {
	printLines()
	fmt.Println()
	fmt.Println("Welcome to the PDNSAPIReader. Set your user and password:")
	fmt.Println()
	printLines()
	setUser()
	setPassword()
}

func printMainCommandScreen() {

	printLines()
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println()
	fmt.Println("COMMANDS:\nSearchDomain (Retrieves a list of all domains that fit the searchkey)\nShowDomain (Retrieves all Records for a given domain)\nsetTTL (Set Time To Live)\n\nExit (Immediatly exit the program)")
	fmt.Println()
	fmt.Println("CLIENT SETTINGS MANAGEMENT:\nChangeUser (Option to set the user which will be used to authenticate to the API Server)\nChangePassword (Option to set the password that will be used to authenticate to the API Server)\nURLChange (Option to change between experimental and production API Servers)\nURLStatus (Checks which environment it's going to call (Default: Production))")
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
	case "changeuser", "cu":
		setUser()
	case "changepassword", "cp":
		setPassword()
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
	confirmedSearchKey := confirm(searchKeyInput)
	fmt.Println()
	printLines()

	preparedForm := SeDForm{
		Action:    "searchdomain",
		Searchkey: confirmedSearchKey,
		User:      user,
		Password:  password,
	}
	readyForm, _ := json.Marshal(preparedForm)
	data := sendPostRequest(readyForm)
	fmt.Println(string(data))
}

func showDomain() {
	var respForm respShDForm

	printLines()
	fmt.Println()
	fmt.Print("What is your domain? ")
	var searchDomainInput string
	fmt.Scanln(&searchDomainInput)
	confirmedDomain := confirm(searchDomainInput)
	fmt.Println()
	printLines()

	preparedForm := ShDForm{
		Action:   "showdomain",
		Domain:   confirmedDomain,
		User:     user,
		Password: password,
	}
	readyForm, _ := json.Marshal(preparedForm)
	data := sendPostRequest(readyForm)
	err := json.Unmarshal(data, &respForm)
	if err != nil {
		log.Println("Unable to unmarshal", err)
	}
	for _, x := range respForm.Domain.Records {
		fmt.Println(x)
	}
	fmt.Println()
	pressAny()
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
}

func urlStatus() {
	printLines()
	fmt.Println()
	fmt.Println("Current URL:", selectedURL)
	fmt.Println()
	printLines()
}

func setUser() {
	fmt.Print("Enter your user: ")
	fmt.Scanln(&user)
}

func setPassword() {
	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)
}

func sendPostRequest(form []byte) []byte {

	resp, err := http.Post(selectedURL, "application/json", bytes.NewBuffer(form))
	if err != nil {
		log.Println("Error encountered:", err)
	}
	rawData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error encountered:", err)
	}
	defer resp.Body.Close()
	return rawData
}

func printLines() {
	fmt.Println(strings.Join([]string{strings.Repeat("=", 80)}, " "))
}

func pressAny() {
	fmt.Println("Press any key to continue...")
	fmt.Scanln()
}
