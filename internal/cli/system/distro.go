package system

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Distro() string {
	currentOS := runtime.GOOS
	switch currentOS {
	case "windows":
		return "windows"
	case "darwin":
		return "macos"
	case "linux":
		path := "/data/data/com.termux/"
		prefix := os.Getenv("PREFIX")
		var prefixdir bool
		if _, err := os.Stat(prefix); os.IsNotExist(err) {
			prefixdir = false
		} else {
			prefixdir = true
		}
		var cmds bool
		cmd := exec.Command("command", "-v", "pkg")
		err := cmd.Run()
		if err != nil {
			cmds = false
		} else {
			cmds = true
		}
		if PathExists(path) && prefix != "" && prefixdir && cmds {
			return "termux"
		}
		return "linux"
	default:
		return "unknown"
	}
}

func DistroID() string {
	files := []string{"/etc/os-release", "/etc/lsb-release", "/etc/redhat-release"}
	for _, file := range files {
		if PathExists(file) {
			f, err := os.Open(file)
			if err != nil {
				return ""
			}
			defer f.Close()
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, "ID=") {
					return strings.Split(line, "=")[1]
				}
			}
		}
	}
	return ""
}
