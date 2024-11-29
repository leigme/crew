package cmd

import (
	"os"
	"path/filepath"
)

const work_dir string = ".crew"

func configPath() string {
	home, err := os.UserHomeDir()
	if nil != err {
		return "."
	}
	return filepath.Join(home, work_dir)
}
