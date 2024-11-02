package pkgman

import (
	"fmt"
	"os"
	"os/exec"
)

func AUR(option string, pkgs string) {
	pacman, err := exec.LookPath("yay")
	if err != nil {
		fmt.Println("yay is not installed")
		os.Exit(1)
	}
	if pkgs == "" {
		fmt.Println("No package specified")
		os.Exit(1)
	}
	var cmd *exec.Cmd
	switch option {
	case "auri":
		cmd = exec.Command(pacman, "-S", pkgs)
	case "aurinc":
		cmd = exec.Command(pacman, "-S", "--noconfirm", pkgs)
	case "auru":
		cmd = exec.Command(pacman, "-Sy", pkgs)
	case "auruu":
		cmd = exec.Command(pacman, "-Syu", pkgs)
	case "aurs":
		cmd = exec.Command(pacman, "-Ss", pkgs)
	case "aurr":
		cmd = exec.Command(pacman, "-Runscd", pkgs)
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
