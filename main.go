package main

import (
	"context"
	"git-hotspot/hotspot"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get current working directory: %v", err)
	}

	if _, err = os.Stat(dir + "/.git"); os.IsNotExist(err) {
		log.Fatal(".git directory does not exist in current directory")
	}

	if err = hotspot.Run(context.Background(), hotspot.NewGitCommands(), dir, 15); err != nil {
		log.Fatal(err)
	}
}
