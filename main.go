package main

import (
	"fmt"
	"log"

	"github.com/cyberis/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v\n", err)
	}

	err = cfg.SetUser("cyberis")
	if err != nil {
		log.Fatalf("Error setting user: %v\n", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error re-reading config: %v\n", err)
	}
	fmt.Printf("%#v\n", cfg)
}
