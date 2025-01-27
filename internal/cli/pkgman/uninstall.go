package pkgman

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dotcreep/gedebox/internal/cli/system"
	"github.com/dotcreep/gedebox/internal/utils"
)

func Uninstall(pkgs string) {
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
		cmd = exec.Command("pkg", "uninstall", "-y", pkgs)
	case "macos":
		switch macpacman {
		case "brew":
			cmd = exec.Command("brew", "uninstall", pkgs)
		case "port":
			cmd = exec.Command("port", "uninstall", pkgs)
		default:
			if err := utils.OpError(macpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "windows":
		switch winpacman {
		case "choco":
			cmd = exec.Command("choco", "uninstall", "-y", pkgs)
		case "winget":
			cmd = exec.Command("winget", "uninstall", "--accept-source-agreements", "--accept-package-agreements", pkgs)
		case "scoop":
			cmd = exec.Command("scoop", "uninstall", pkgs)
		default:
			if err := utils.OpError(winpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "linux":
		switch pacman {
		case "apt-get", "dnf", "yum", "zypper", "pkgin": // "apt",
			if su == "" {
				cmd = exec.Command(pacman, "remove", "-y", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "remove", "-y", pkgs)
			}
		case "pkg":
			if su == "" {
				cmd = exec.Command(pacman, "delete", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "delete", pkgs)
			}
		case "pkg_add":
			if su == "" {
				cmd = exec.Command("pkg_delete", pkgs)
			} else {
				cmd = exec.Command(su, "pkg_delete", pkgs)
			}
		case "pacman":
			if su == "" {
				cmd = exec.Command(pacman, "-R", "--noconfirm", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "-R", "--noconfirm", pkgs)
			}
		case "xbps-install":
			if su == "" {
				cmd = exec.Command("xbps-remove", "-y", pkgs)
			} else {
				cmd = exec.Command(su, "xbps-remove", "-y", pkgs)
			}
		case "apk":
			if su == "" {
				cmd = exec.Command(pacman, "del", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "del", pkgs)
			}
		case "emerge":
			if su == "" {
				cmd = exec.Command(pacman, "--deselect", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "--deselect", pkgs)
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
