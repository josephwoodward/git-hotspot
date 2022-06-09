package hotspot

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type File struct {
	File string
	Log  []Log
}

type Log struct {
	Date string
}

func Run(dir string) error {
	list, err := exec.Command("git", "-C", dir, "ls-files").Output()
	if err != nil {
		return err
	}

	//var items []File
	files := bytes.Split(list, []byte("\n"))
	//files := strings.Split(list, "\n")

	_ = files

	//c := make(chan File)
	//for _, file := range files {
	//	items = append(items, File{File: file})
	//	//getLog(c, dir+"/"+file)
	//}

	return nil
}

func getLog(c chan File, path string) {
	cmd := exec.Command("git", "log", "--pretty=%ad", "--date=short", path)
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
	//c <- &File{Log: }
}
