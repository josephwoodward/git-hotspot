package hotspot

import (
	"bytes"
	"context"
	"log"
	"os/exec"
	"sort"
	"sync"
)

func Run(ctx context.Context, dir string) error {
	list, err := exec.Command("git", "-C", dir, "ls-files").Output()
	if err != nil {
		return err
	}

	results := make(chan file)

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
			modifications(ctx, string(f), results)
		}(f)
	}

	var total []file
	go func() {
		for f := range results {
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

	//for i := range s {
	//	log.Printf("%s: %v", total[i].path, len(total[i].dates))
	//}
	for _, d := range s {
		//fmt.Printf("%+v\n", *d.dates)
		log.Printf("%s: %v", d.path, len(d.dates))
	}

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

	results <- file{path: path, dates: dates}
	return
}
