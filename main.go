package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type File struct {
	File string
	Log  []Log
}

type Log struct {
	Date string
}

func main() {

	workingDir := getWorkingDir()
	cmd := exec.Command("git", "-C", workingDir, "ls-files")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	var items []File
	files := strings.Split(out.String(), "\n")
	for _, file := range files {
		items = append(items, File{File: file})
		getLog(workingDir + "/" + file)
	}

}

func getLog(file string) {
	cmd := exec.Command("git", "log", "--pretty=%ad", "--date=short", file)
	var stOut bytes.Buffer
	var stErr bytes.Buffer
	cmd.Stdout = &stOut
	cmd.Stderr = &stErr
	err := cmd.Run()
	if err != nil {
		log.Fatalf(stErr.String(), "There was an error running the git command: ", err)
		os.Exit(1)
	}

	fmt.Println(stOut.String())
}

func getWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}
