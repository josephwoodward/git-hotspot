package main

import (
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get current working directory: %w", err)
	}

	if err := hotspot.Run(dir); err != nil {
		log.Fatal(err)
	}
}
