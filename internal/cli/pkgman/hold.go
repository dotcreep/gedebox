package pkgman

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dotcreep/gedebox/internal/cli/system"
	"github.com/dotcreep/gedebox/internal/utils"
)

func Hold(pkgs string) {
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
		cmd = exec.Command("pkg", "hold", pkgs)
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
				cmd = exec.Command("apt-mark", "hold", pkgs)
			} else {
				cmd = exec.Command(su, "apt-mark", "hold", pkgs)
			}
		case "zypper":
			if su == "" {
				cmd = exec.Command(pacman, "addlock", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "addlock", pkgs)
			}
		case "pacman":
			if su == "" {
				cmd = exec.Command(pacman, "-D", "--asexplicit", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "-D", "--asexplicit", pkgs)
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
