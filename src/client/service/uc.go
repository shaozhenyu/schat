package service

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Login(args []string) {
	//TODO
	fmt.Println("login", args)
	fmt.Println("Please input user name: ")
	read := bufio.NewReader(os.Stdin)
	username, _, err := read.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(username))
	fmt.Println("Please input password: ")
	passwd, _, err := read.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(passwd))
	//mongo find
}

func Register(args []string) {
	//TODO
}
