package main

import (
	"context"
	"git-hotspot/hotspot"
	"log"
	"os"
	"os/exec"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get current working directory: %v", err)
	}

	if _, err = os.Stat(dir + "/.git"); os.IsNotExist(err) {
		log.Fatal(".git directory does not exist in current directory")
	}

	cli := git{}
	if err = hotspot.Run(context.Background(), cli, 15); err != nil {
		log.Fatal(err)
	}
}

// git satisfies the hotspot.GitCommands interface
type git struct {
	dir string
}

// Config returns paths for hotspot to ignore
func (c git) Config() ([]byte, error) {
	return exec.Command("git", "config", "--get-all", "hotspot.path").Output()
}

// Files lists all files in the repository
func (c git) Files() ([]byte, error) {
	return exec.Command("git", "-C", c.dir, "ls-files").Output()
}
