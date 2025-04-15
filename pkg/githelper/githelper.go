package githelper

import (
	"bytes"
	"os/exec"
	"strings"
)

type GitHelper struct {
	conf *GitHelperConfig
}

func NewGitHelper(conf *GitHelperConfig) *GitHelper {
	return &GitHelper{
		conf: conf,
	}
}

func (h *GitHelper) GetMergeBase() string {
	cmd := exec.Command("git", "merge-base", branch, "HEAD")
	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(buf.String())
}

type GitHelperConfig struct {
}
