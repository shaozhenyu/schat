package main

import (
	"fmt"
	"lib/odm"
)

func main() {
	fmt.Println("server")
	odm.New("127.0.0.1:27017", "uc")
}
