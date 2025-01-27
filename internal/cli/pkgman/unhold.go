package pkgman

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dotcreep/gedebox/internal/cli/system"
	"github.com/dotcreep/gedebox/internal/utils"
)

func Unhold(pkgs string) {
	if pkgs == "" {
		fmt.Println("No package specified")
		os.Exit(1)
	}
	distro := system.Distro()
	distro_id := system.DistroID()
	pacman := system.GetLinuxPackageManager(distro_id)
	macpacman := system.GetMacPackageManager()
	winpacman := system.GetWindowsPackageManager()
	su := system.User()
	var cmd *exec.Cmd
	switch distro {
	case "termux":
		cmd = exec.Command("pkg", "unhold", pkgs)
	case "macos":
		if err := utils.OpError(macpacman); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "windows":
		if err := utils.OpError(winpacman); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "linux":
		switch pacman {
		case "apt", "apt-get":
			if su == "" {
				cmd = exec.Command("apt-mark", "unhold", pkgs)
			} else {
				cmd = exec.Command(su, "apt-mark", "unhold", pkgs)
			}
		case "zypper":
			if su == "" {
				cmd = exec.Command(pacman, "removelock", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "removelock", pkgs)
			}
		case "pacman":
			if su == "" {
				cmd = exec.Command(pacman, "-D", "--asdeps", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "-D", "--asdeps", pkgs)
			}
		case "xbps-install":
			if su == "" {
				cmd = exec.Command("xbps-query", "-Rs", pkgs)
			} else {
				cmd = exec.Command(su, "xbps-query", "-Rs", pkgs)
			}
		case "emerge":
			if su == "" {
				cmd = exec.Command("bash", "-c", fmt.Sprintf("sed -i \"/%s/d\" /etc/portage/package.mask", pkgs))
			} else {
				cmd = exec.Command(su, "bash", "-c", fmt.Sprintf("sed -i \"/%s/d\" /etc/portage/package.mask", pkgs))
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
