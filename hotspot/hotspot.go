package hotspot

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"sort"
	"sync"

	"github.com/rodaine/table"
)

func Run(ctx context.Context, git GitCommands, limit int) error {
	_, err := parseConfig(git)

	// List Files
	result, err := git.Files()
	if err != nil {
		return err
	}

	results := make(chan file)

	// Iterate over all Files in repo
	files := bytes.Split(result, []byte("\n"))
	var wg sync.WaitGroup
	for _, f := range files {
		if len(f) == 0 {
			continue
		}

		wg.Add(1)
		go func(f []byte) {
			defer wg.Done()
			modifications(ctx, string(f), results)
		}(f)
	}

	var total []file
	go func() {
		for f := range results {
			if len(f.dates) == 0 {
				continue
			}

			total = append(total, f)
		}
	}()

	wg.Wait()
	close(results)

	// Sort results
	s := make(dataSlice, 0, len(total))
	for k, _ := range total {
		s = append(s, &total[k])
	}

	sort.Sort(s)

	// Print results
	tbl := table.
		New("File", "# Modifications").
		WithWriter(os.Stdout)

	i := 0
	for _, d := range s {
		if i == limit {
			break
		}
		tbl.AddRow(d.path, len(d.dates))
		i++
	}

	tbl.Print()

	return nil
}

// modifications reports the number of edits
func modifications(_ context.Context, path string, results chan<- file) {
	logs, err := exec.Command("git", "log", "--pretty=%ad", "--date=short", path).Output()
	if err != nil {
		results <- file{path: path, err: err, dates: nil}
	}

	l := bytes.Split(logs, []byte("\n"))

	var dates []string
	for _, v := range l {
		if len(v) > 0 {
			dates = append(dates, string(v))
		}
	}

	if len(dates) > 0 {
		results <- file{path: path, dates: dates}
	}

	return
}

func parseConfig(cli GitCommands) (map[string]struct{}, error) {
	result, err := cli.Config()
	if err != nil {
		return nil, err
	}

	ignores := bytes.Split(result, []byte("\n"))
	paths := map[string]struct{}{}

	for _, v := range ignores {
		k := string(v)
		paths[k] = struct{}{}
	}

	return paths, err
}
