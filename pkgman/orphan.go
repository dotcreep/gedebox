package pkgman

import (
	"fmt"
	"gedebox/handler"
	"gedebox/system"
	"os"
	"os/exec"
)

func Orphan() {
	distro := system.Distro()
	distro_id := system.DistroID()
	pacman := system.GetLinuxPackageManager(distro_id)
	macpacman := system.GetMacPackageManager()
	winpacman := system.GetWindowsPackageManager()
	su := system.User()
	var cmd *exec.Cmd
	switch distro {
	case "termux":
		cmd = exec.Command("pkg", "autoremove")
	case "macos":
		switch macpacman {
		case "brew":
			cmd = exec.Command("brew", "cleanup")
		default:
			if err := handler.OpError(macpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "windows":
		switch winpacman {
		case "choco":
			cmd = exec.Command("choco", "clean")
		case "winget":
			cmd = exec.Command("winget", "clean")
		case "scoop":
			cmd = exec.Command("scoop", "cleanup")
		default:
			if err := handler.OpError(winpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "linux":
		switch pacman {
		case "apt", "apt-get", "dnf", "yum":
			if su == "" {
				cmd = exec.Command(pacman, "autoremove")
			} else {
				cmd = exec.Command(su, pacman, "autoremove")
			}
		case "zypper":
			if su == "" {
				cmd = exec.Command(pacman, "clean")
			} else {
				cmd = exec.Command(su, pacman, "clean")
			}
		case "pacman":
			if su == "" {
				cmd = exec.Command(pacman, "-Rns", "$(pacman -Qdtq)")
			} else {
				cmd = exec.Command(su, pacman, "-Rns", "$(pacman -Qdtq)")
			}
		case "xbps-install":
			if su == "" {
				cmd = exec.Command("xbps-remove", "-o")
			} else {
				cmd = exec.Command(su, "xbps-remove", "-o")
			}
		case "apk":
			if su == "" {
				cmd = exec.Command(pacman, "orphans")
			} else {
				cmd = exec.Command(su, pacman, "orphans")
			}
		case "emerge":
			cmd = exec.Command(su, pacman, "--depclean")
		default:
			if err := handler.OpError(pacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	default:
		if err := handler.DistError(distro); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
