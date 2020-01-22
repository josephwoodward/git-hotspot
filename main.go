package main

import (
	"bytes"
	"fmt"
	"github.com/kardianos/osext"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("git", "-C", getExecutingFolder(), "ls-files")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	files := strings.Split(out.String(), "\n")
	for _, file := range files {
		//getLog(file)
		runBasicExample(file)
	}

}

func runBasicExample(file string) {
	cmdOut := exec.Command("git", "log", "/Users/joseph.woodward/Documents/git/AccountWeb/tests/JustEat.AccountWeb.Tests.Integration/HomepageTests.cs")
	var stOut bytes.Buffer
	cmdOut.Stdout = &stOut
	cmdOut.Run()

	fmt.Println(stOut.String())

}

func getLog(file string) {
	//cmd := exec.Command("git", "log", "~/Documents/git/AccountWeb/tests/JustEat.AccountWeb.Tests.Integration/HomepageTests.cs", "--pretty='format: %ad'", "--date=short", "-- \"$1\"")
	cmd := exec.Command("git", "log", "/Users/joseph.woodward/Documents/git/AccountWeb/tests/JustEat.AccountWeb.Tests.Integration/HomepageTests.cs")
	var stOut bytes.Buffer
	var stErr bytes.Buffer
	cmd.Stdout = &stOut
	cmd.Stderr = &stErr
	err := cmd.Run()
	if err != nil {
		log.Fatalf(stErr.String(), "There was an error running the git command: ", err)
		os.Exit(1)	}

	fmt.Println(stOut.String())
}


func getExecutingFolder() string {
	_, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}

	return "/Users/joseph.woodward/Documents/git/AccountWeb"
}