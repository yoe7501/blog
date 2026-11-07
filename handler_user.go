package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd Command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("arguments are empty")
	}
	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Println("User set!")
	return nil
}
