package system

import (
	"os"
	"os/exec"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func Where(command string) string {
	path, err := exec.LookPath(command)
	if err != nil {
		return ""
	}
	return path
}
