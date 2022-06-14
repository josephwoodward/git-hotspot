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

	if err = hotspot.Run(context.Background(), dir, 0); err != nil {
		t.Fatal(err)
	}
}
