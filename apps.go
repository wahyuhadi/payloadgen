package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	payload = flag.String("payload", "", "Type payload  java, python, php")
	command = flag.String("command", "", "Type command like , cat /etc/passwd")
)

func main() {
	flag.Parse()
	if *payload == "java" {
		readJava()
	}

	fmt.Println("Done ..")
}

func readJava() {
	flag.Parse()
	data, err := ioutil.ReadFile("DO_NOT_DELETE_THIS_FOLDER/javapayload.file")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	newContents := strings.Replace(string(data), "codeinjection", *command, -1)

	fmt.Println(newContents)
}
