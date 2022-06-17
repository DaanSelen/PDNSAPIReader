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

	urlArray[0] = cfg.Section("production").Key("server").String()
	urlArray[1] = cfg.Section("test").Key("server").String()
	user[0] = cfg.Section("production").Key("user").String()
	user[1] = cfg.Section("test").Key("user").String()
	password[0] = cfg.Section("production").Key("password").String()
	password[1] = cfg.Section("test").Key("password").String()
}

func main() {
	go getIniData()
	initFlags()

	checkEnv()
	fmt.Println(*actionPointer)
	fmt.Println(*domainPointer)
	fmt.Println(*valuePointer)
	fmt.Println(selectedURL)
}

func checkEnv() {
	switch *envPointer {
	case "prod":
		selectedURL = urlArray[0]
	case "test":
		selectedURL = urlArray[1]
	default:
		log.Println("Invalid flag entry:", *envPointer)
	}
}
