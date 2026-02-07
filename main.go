package main

import (
	"fmt"

	"github.com/cyberis/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return
	}
	err = config.SetUser("cyberis")
	if err != nil {
		fmt.Printf("Error setting user: %v\n", err)
		return
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return
	}
	fmt.Printf("%#v\n", cfg)
}
