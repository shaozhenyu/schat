package main

import (
	"client/service"
	"ember/cli"
	"os"
)

func main() {
	cmds := cli.NewCmds()
	Reg(cmds)

	args := os.Args[1:]
	if len(args) == 0 {
		cmds.Help(true)
	} else {
		cmds.Run(args)
	}
}

func Reg(cmds *cli.Cmds) {
	cmds.Reg("login", "user login", service.Login)
	cmds.Reg("register", "register new user", service.Register)
}
