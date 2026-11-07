package main

import "errors"

type Command struct {
	name      string
	arguments []string
}

type Commands struct {
	commandToHandler map[string]func(*state, Command) error
}

func (c *Commands) run(s *state, cmd Command) error {
	handler, ok := c.commandToHandler[cmd.name]
	if !ok {
		return errors.New("command doesn't exist")
	}

	err := handler(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *Commands) register(name string, f func(*state, Command) error) {
	c.commandToHandler[name] = f
}
