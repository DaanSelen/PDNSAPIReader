package main

import (
	"flag"
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

var (
	envPointer    *string
	actionPointer *string
	domainPointer *string
	valuePointer  *string

	selectedURL string
	urlArray    [2]string
	user        [2]string
	password    [2]string
)

func initFlags() {
	envPointer = flag.String("env", "prod", "setting the environment to call to")
	actionPointer = flag.String("a", "", "setting the action")
	domainPointer = flag.String("d", "", "setting the domain to use")
	valuePointer = flag.String("v", "", "setting the environment to call to")
	flag.Parse()
}

func getIniData() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		log.Println("Unable to load the .ini file, error:", err)
	}

	urlArray[0] = cfg.Section("prod").Key("server").String()
	urlArray[1] = cfg.Section("test").Key("server").String()
	user[0] = cfg.Section("prod").Key("user").String()
	user[1] = cfg.Section("test").Key("user").String()
	password[0] = cfg.Section("prod").Key("password").String()
	password[1] = cfg.Section("test").Key("password").String()
}

func main() {
	getIniData()
	initFlags()

	if checkEnv() {
		fmt.Println(selectedURL)
		initAction()
	}
}

func checkEnv() bool {
	switch *envPointer {
	case "prod":
		selectedURL = urlArray[0]
		return true
	case "test":
		selectedURL = urlArray[1]
		return true
	default:
		log.Println("Invalid flag entry:", *envPointer)
		return false
	}
}

func initAction() {
	if *actionPointer != "" && *domainPointer != "" && *valuePointer != "" {
		if *actionPointer == "setttl" {
			log.Println("Selected setttl with", *domainPointer, "and", *valuePointer)
		} else {
			log.Println("ERROR: Missing or incorrect flags")
		}
	} else if *actionPointer != "" && *domainPointer != "" {
		switch *actionPointer {
		case "searchdomain", "sed":
			log.Println("selected searchdomain with", *domainPointer)
		case "showdomain", "shd":
			log.Println("selected showdomain with", *domainPointer)
		default:
			log.Println("ERROR: Missing or incorrect flags")
		}
	} else {
		log.Println("ERROR: Missing or incorrect flags")
	}
}
