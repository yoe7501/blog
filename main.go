package main

import (
	"fmt"
	"os"

	"github.com/yoe7501/blog/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	s := state{
		config: &c,
	}
	cmds := Commands{
		commandToHandler: make(map[string]func(*state, Command) error),
	}
	cmds.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage [command] <arguments>")
		os.Exit(1)
	}
	cmd := Command{
		name:      args[1],
		arguments: args[2:],
	}
	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
