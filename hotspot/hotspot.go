package hotspot

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type File struct {
	File string
	Log  []Log
	Logs []byte
}

type Log struct {
	Date string
}

func Run(dir string) error {
	list, err := exec.Command("git", "-C", dir, "ls-files").Output()
	if err != nil {
		return err
	}

	ctx := context.Background()

	// Iterate over all files in repo
	files := bytes.Split(list, []byte("\n"))
	done := make(chan struct{})
	i := 0
	for f := range modifications(ctx, done, files) {
		if i == 4 {
			done <- struct{}{}
		}
		fmt.Println(len(f.Logs))
		i++
	}

	//c := make(chan File)
	//for _, file := range files {
	//	items = append(items, File{File: file})
	//	//getLog(c, dir+"/"+file)
	//}

	return nil
}

// modifications reports the number of edits
func modifications(ctx context.Context, done <-chan struct{}, files [][]byte) <-chan File {
	results := make(chan File)

	for _, f := range files {
		if len(f) == 0 {
			continue
		}

		go func(done <-chan struct{}, fi []byte) {
			f := string(fi)
			logs, err := exec.Command("git", "log", "--pretty=%ad", "--date=short", f).Output()
			if err != nil {
				os.Exit(1)
			}

			time.Sleep(1 * time.Second)

			r := File{
				File: f,
				Logs: logs,
			}
			select {
			case <-done:
				return
			case results <- r:
			}
		}(done, f)

	}

	return results
}
