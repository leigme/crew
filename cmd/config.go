package cmd

import (
	"os"
	"path/filepath"
)

const workDir string = ".crew"

func configPath() string {
	home, err := os.UserHomeDir()
	if nil != err {
		return "."
	}
	return filepath.Join(home, workDir)
}
