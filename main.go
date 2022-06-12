package main

import (
	"context"
	"fmt"
	"git-hotspot/hotspot"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get current working directory: %v", err)
	}

	if err := hotspot.Run(context.Background(), dir); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
