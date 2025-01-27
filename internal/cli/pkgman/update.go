package pkgman

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dotcreep/gedebox/internal/cli/system"
	"github.com/dotcreep/gedebox/internal/utils"
)

func Update() {
	distro := system.Distro()
	distro_id := system.DistroID()
	pacman := system.GetLinuxPackageManager(distro_id)
	winpacman := system.GetWindowsPackageManager()
	macpacman := system.GetMacPackageManager()
	su := system.User()
	var cmd *exec.Cmd
	switch distro {
	case "termux":
		cmd = exec.Command("pkg", "update")
	case "macos":
		switch macpacman {
		case "brew":
			cmd = exec.Command(macpacman, "update")
		case "port":
			cmd = exec.Command(macpacman, "selfupdate")
		default:
			if err := utils.OpError(macpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "windows":
		switch winpacman {
		case "choco":
			cmd = exec.Command("choco", "upgrade")
		case "winget":
			cmd = exec.Command("winget", "upgrade --all")
		case "scoop":
			cmd = exec.Command("scoop", "update")
		default:
			if err := utils.OpError(winpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "linux":
		switch pacman {
		case "apt-get", "pkg", "pkgin": // "apt",
			if su == "" {
				cmd = exec.Command(pacman, "update")
			} else {
				cmd = exec.Command(su, pacman, "update")
			}
		case "dnf", "yum":
			if su == "" {
				cmd = exec.Command(pacman, "check-update")
			} else {
				cmd = exec.Command(su, pacman, "check-update")
			}
		case "zypper":
			if su == "" {
				cmd = exec.Command(pacman, "refresh")
			} else {
				cmd = exec.Command(su, pacman, "refresh")
			}
		case "pacman", "xbps-install":
			if su == "" {
				cmd = exec.Command(pacman, "-Sy")
			} else {
				cmd = exec.Command(su, pacman, "-Sy")
			}
		case "apk":
			if su == "" {
				cmd = exec.Command(pacman, "update")
			} else {
				cmd = exec.Command(su, pacman, "update")
			}
		case "emerge":
			if su == "" {
				cmd = exec.Command(pacman, "sync")
			} else {
				cmd = exec.Command(su, pacman, "sync")
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
