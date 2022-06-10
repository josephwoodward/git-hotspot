package main_test

import (
	"git-hotspot/hotspot"
	"os"
	"testing"
)

func TestHotspot(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed getting current working directory: %v", err)
	}

	if err := hotspot.Run(dir); err != nil {
		t.Fatal(err)
	}
}
