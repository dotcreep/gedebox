package pkgman

import (
	"fmt"
	"gedebox/handler"
	"gedebox/system"
	"os"
	"os/exec"
)

func Upgrade() {
	distro := system.Distro()
	distro_id := system.DistroID()
	pacman := system.GetLinuxPackageManager(distro_id)
	winpacman := system.GetWindowsPackageManager()
	macpacman := system.GetMacPackageManager()
	su := system.User()
	var cmd *exec.Cmd
	switch distro {
	case "termux":
		cmd = exec.Command("pkg", "upgrade")
	case "macos":
		switch macpacman {
		case "brew":
			cmd = exec.Command("brew", "update")
		case "port":
			cmd = exec.Command("port", "upgrade")
		default:
			if err := handler.OpError(macpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "windows":
		switch winpacman {
		case "choco":
			cmd = exec.Command("choco", "upgrade", "all", "-y")
		case "winget":
			cmd = exec.Command("winget", "upgrade", "--all")
		case "scoop":
			cmd = exec.Command("scoop", "update", "*")
		default:
			if err := handler.OpError(winpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "linux":
		switch pacman {
		case "apt", "apt-get", "pkg":
			if su == "" {
				cmd = exec.Command(pacman, "upgrade", "-y")
			} else {
				cmd = exec.Command(su, pacman, "upgrade", "-y")
			}
		case "dnf", "yum":
			if su == "" {
				cmd = exec.Command(pacman, "upgrade", "-y")
			} else {
				cmd = exec.Command(su, pacman, "upgrade", "-y")
			}
		case "zypper":
			if su == "" {
				cmd = exec.Command(pacman, "--non-interactive", "update")
			} else {
				cmd = exec.Command(su, pacman, "--non-interactive", "update")
			}
		case "pacman":
			if su == "" {
				cmd = exec.Command(pacman, "-Syu", "--noconfirm")
			} else {
				cmd = exec.Command(su, pacman, "-Syu", "--noconfirm")
			}
		case "xbps-install":
			if su == "" {
				cmd = exec.Command(pacman, "-Syu")
			} else {
				cmd = exec.Command(su, pacman, "-Syu")
			}
		case "apk":
			if su == "" {
				cmd = exec.Command(pacman, "upgrade", "--no-cache", "--update")
			} else {
				cmd = exec.Command(su, pacman, "upgrade", "--no-cache", "--update")
			}
		default:
			if err := handler.OpError(pacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	default:
		if err := handler.OpError(distro); err != nil {
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
