package pkgman

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dotcreep/gedebox/internal/cli/system"
)

func SNAP(option string, pkgs string) {
	snap, err := exec.LookPath("snap")
	if err != nil {
		fmt.Println("snap is not installed")
		os.Exit(1)
	}
	if pkgs == "" {
		fmt.Println("No package specified")
		os.Exit(1)
	}
	su := system.User()
	var cmd *exec.Cmd
	switch option {
	case "snapi":
		if su == "" {
			cmd = exec.Command(snap, "install", pkgs)
		} else {
			cmd = exec.Command(su, snap, "install", pkgs)
		}
	case "snapu":
		if su == "" {
			cmd = exec.Command(snap, "refresh", pkgs)
		} else {
			cmd = exec.Command(su, snap, "refresh", pkgs)
		}
	case "snapv":
		if su == "" {
			cmd = exec.Command(snap, "revert", pkgs)
		} else {
			cmd = exec.Command(su, snap, "revert", pkgs)
		}
	case "snaps":
		if su == "" {
			cmd = exec.Command(snap, "find", pkgs)
		} else {
			cmd = exec.Command(su, snap, "find", pkgs)
		}
	case "snapl":
		if su == "" {
			cmd = exec.Command(snap, "list", pkgs)
		} else {
			cmd = exec.Command(su, snap, "list", pkgs)
		}
	case "snapla":
		if su == "" {
			cmd = exec.Command(snap, "list", "--all", pkgs)
		} else {
			cmd = exec.Command(su, snap, "list", "--all", pkgs)
		}
	case "snapon":
		if su == "" {
			cmd = exec.Command(snap, "enable", pkgs)
		} else {
			cmd = exec.Command(su, snap, "enable", pkgs)
		}
	case "snapoff":
		if su == "" {
			cmd = exec.Command(snap, "disable", pkgs)
		} else {
			cmd = exec.Command(su, snap, "disable", pkgs)
		}
	case "snapr":
		if su == "" {
			cmd = exec.Command(snap, "remove", pkgs)
		} else {
			cmd = exec.Command(su, snap, "remove", pkgs)
		}
	default:
		fmt.Println("Error: invalid option")
		os.Exit(1)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
