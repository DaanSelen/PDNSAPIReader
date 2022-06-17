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

	"gopkg.in/ini.v1"
)

var (
	selectedIndex = 0
	selectedURL   = urlArray[selectedIndex]
	urlArray      [2]string
	user          [2]string
	password      [2]string
)

func main() {
	go getIniData()
	printMainStartScreen()
	printMainCommandScreen()
}

func getIniData() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		log.Println("Unable to load the .ini file, error:", err)
	}

	go func() {
		urlArray[0] = cfg.Section("production").Key("server").String()
		urlArray[1] = cfg.Section("test").Key("server").String()
	}()
	go func() {
		user[0] = cfg.Section("production").Key("user").String()
		user[1] = cfg.Section("test").Key("user").String()
	}()
	go func() {
		password[0] = cfg.Section("production").Key("password").String()
		password[1] = cfg.Section("test").Key("password").String()
	}()
}

func printMainStartScreen() {
	printLines()
	fmt.Println()
	fmt.Println("Welcome to the PDNSAPIReader.")
	fmt.Println()
	printLines()
}

func printMainCommandScreen() {

	printLines()
	fmt.Println()
	fmt.Println("Current API URL:", selectedURL)
	fmt.Println("Available commands:")
	fmt.Println()
	fmt.Println("COMMANDS:\nSearchDomain (Retrieves a list of all domains that fit the searchkey)\nShowDomain (Retrieves all Records for a given domain)\nsetTTL (Set Time To Live)\n\nExit (Immediatly exit the program)")
	fmt.Println()
	fmt.Println("CLIENT SETTINGS MANAGEMENT:\nURLChange (Option to change between experimental and production API Servers)")
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
	}
	printMainCommandScreen()
}

func searchDomain() {
	var respForm respSeDForm

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
		User:      user[selectedIndex],
		Password:  password[selectedIndex],
	}
	readyForm, _ := json.Marshal(preparedForm)
	data := sendPostRequest(readyForm)
	err := json.Unmarshal(data, &respForm)
	if err != nil {
		log.Println("Unable to unmarshal", err)
	}
	for _, y := range respForm.Domains {
		fmt.Println(y)
	}
	fmt.Println()
	pressAny()
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
		User:     user[selectedIndex],
		Password: password[selectedIndex],
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
	var respForm respTTLForm

	printLines()
	fmt.Println()
	fmt.Print("What is the domain you wish to set the TTL for? ")
	var ttlDomainInput string
	fmt.Scanln(&ttlDomainInput)
	confirmedTTLDomain := confirm(ttlDomainInput)
	fmt.Print("What is the TTL you wish to set for the given domain? ")
	var desiredTTLInput string
	fmt.Scanln(&desiredTTLInput)
	confirmedDesiredTTL := confirm(desiredTTLInput)
	fmt.Print("What is the reason (optional)? ")
	var reasonInput string
	fmt.Scanln(&reasonInput)

	preparedForm := STForm{
		Action:     "setTTL",
		Domainname: confirmedTTLDomain,
		TTL:        confirmedDesiredTTL,
		Reason:     reasonInput,
		User:       user[selectedIndex],
		Password:   password[selectedIndex],
	}
	readyForm, _ := json.Marshal(preparedForm)
	data := sendPostRequest(readyForm)
	err := json.Unmarshal(data, &respForm)
	if err != nil {
		log.Println("Unable to unmarshal", err)
	}
	fmt.Println(respForm.Message)
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
