package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("username argument is required")
	}
	userName := cmd.args[0]
	err := s.Config.SetUser(userName)
	if err != nil {
		return err
	}
	fmt.Println("User set to:", userName)
	return nil
}
