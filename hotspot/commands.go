package hotspot

import "os/exec"

type git struct {
}

// NewGitCommands creates a default git cli
func NewGitCommands() git {
	return git{}
}

// Config returns paths for hotspot to ignore
func (c git) Config() ([]byte, error) {
	return exec.Command("git", "config", "--get-all", "hotspot.ignorepaths").Output()
}

// Files lists all files in the repository
func (c git) Files(dir string) ([]byte, error) {
	return exec.Command("git", "-C", dir, "ls-files").Output()
}
