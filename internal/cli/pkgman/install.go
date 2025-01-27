package pkgman

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dotcreep/gedebox/internal/cli/system"
	"github.com/dotcreep/gedebox/internal/utils"
)

func Install(pkgs string) {
	if pkgs == "" {
		fmt.Println("No package specified")
		os.Exit(1)
	}
	distro := system.Distro()
	distro_id := system.DistroID()
	pacman := system.GetLinuxPackageManager(distro_id)
	winpacman := system.GetWindowsPackageManager()
	macpacman := system.GetMacPackageManager()
	su := system.User()
	var cmd *exec.Cmd
	switch distro {
	case "termux":
		cmd = exec.Command("pkg", "install", "-y", pkgs)
	case "macos":
		switch macpacman {
		case "brew":
			cmd = exec.Command(macpacman, "install", pkgs)
		case "port":
			cmd = exec.Command(macpacman, "install", pkgs)
		default:
			if err := utils.OpError(macpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "windows":
		switch winpacman {
		case "choco":
			cmd = exec.Command(winpacman, "install", "-y", pkgs)
		case "winget":
			cmd = exec.Command(winpacman, "install", "--silent", "--accept-source-agreements", "--accept-package-agreements", pkgs)
		case "scoop":
			cmd = exec.Command(winpacman, "install", pkgs)
		default:
			if err := utils.OpError(winpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "linux":
		switch pacman {
		case "apt-get", "dnf", "yum", "zypper": // "apt",
			if su == "" {
				cmd = exec.Command(pacman, "install", "-y", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "install", "-y", pkgs)
			}
		case "pkg", "pkgin":
			if su == "" {
				cmd = exec.Command(pacman, "install", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "install", pkgs)
			}
		case "emerge", "pkg_add":
			if su == "" {
				cmd = exec.Command(pacman, pkgs)
			} else {
				cmd = exec.Command(su, pacman, pkgs)
			}
		case "pacman":
			if su == "" {
				cmd = exec.Command(pacman, "-S", "--noconfirm", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "-S", "--noconfirm", pkgs)
			}
		case "xbps-install":
			if su == "" {
				cmd = exec.Command(pacman, "-y", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "-y", pkgs)
			}
		case "apk":
			if su == "" {
				cmd = exec.Command(pacman, "add", "--no-cache", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "add", "--no-cache", pkgs)
			}
		default:
			if err := utils.OpError(pacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	default:
		if err := utils.DistError(distro); err != nil {
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
