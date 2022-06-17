package main_test

import (
	"context"
	"git-hotspot/hotspot"
	"os"
	"testing"
)

func TestHotspot(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed getting current working directory: %v", err)
	}

	git := hotspot.NewGitCommands()
	if err = hotspot.Run(context.Background(), git, dir, 100); err != nil {
		t.Fatal(err)
	}
}
