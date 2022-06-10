package hotspot

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"sync"
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
	var wg sync.WaitGroup

	i := 0
	for f := range modifications(ctx, files, wg) {
		i++
		fmt.Println(len(f.Logs))
		if i == 4 {
			break
		}
	}

	wg.Wait()

	return nil
}

// modifications reports the number of edits
func modifications(ctx context.Context, files [][]byte, wg sync.WaitGroup) <-chan File {
	results := make(chan File)

	for _, v := range files {
		if len(v) == 0 {
			continue
		}

		go func(file []byte) {
			wg.Add(1)
			defer wg.Done()

			f := string(file)
			logs, err := exec.Command("git", "log", "--pretty=%ad", "--date=short", f).Output()
			if err != nil {
				os.Exit(1)
			}

			time.Sleep(1 * time.Second)

			results <- File{
				File: f,
				Logs: logs,
			}
		}(v)
	}

	return results
}
