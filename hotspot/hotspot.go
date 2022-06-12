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
	File  string
	Dates []string
}

func Run(ctx context.Context, dir string) error {
	list, err := exec.Command("git", "-C", dir, "ls-files").Output()
	if err != nil {
		return err
	}

	results := make(chan File)

	// Iterate over all files in repo
	files := bytes.Split(list, []byte("\n"))
	var wg sync.WaitGroup
	for _, f := range files {
		if len(f) == 0 {
			continue
		}

		wg.Add(1)
		go func(f []byte) {
			defer wg.Done()
			modifications(ctx, string(f), results, wg)
		}(f)
	}

	var total []File
	go func() {
		for f := range results {
			total = append(total, f)
		}
	}()

	//sort.Slice(total, func (d, e int) bool {
	//	return a[d].size < a[e].size
	//})

	wg.Wait()
	close(results)

	fmt.Println("Done")

	return nil
}

// modifications reports the number of edits
func modifications(_ context.Context, file string, results chan<- File, wg sync.WaitGroup) {
	logs, err := exec.Command("git", "log", "--pretty=%ad", "--date=short", file).Output()
	if err != nil {
		os.Exit(1)
	}

	time.Sleep(300 * time.Millisecond)

	l := bytes.Split(logs, []byte("\n"))

	var dates []string
	for _, v := range l {
		if len(v) > 0 {
			dates = append(dates, string(v))
		}
	}
	results <- File{
		File:  file,
		Dates: dates,
	}
	return
}
