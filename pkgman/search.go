package pkgman

import (
	"fmt"
	"gedebox/handler"
	"gedebox/system"
	"os"
	"os/exec"
)

func Search(pkgs string) {
	if pkgs == "" {
		fmt.Println("No package specified")
		os.Exit(1)
	}
	distro := system.Distro()
	distro_id := system.DistroID()
	pacman := system.GetLinuxPackageManager(distro_id)
	winpacman := system.GetWindowsPackageManager()
	macpacman := system.GetMacPackageManager()
	var cmd *exec.Cmd
	switch distro {
	case "termux":
		cmd = exec.Command("pkg", "search", pkgs)
	case "macos":
		switch macpacman {
		case "brew", "port":
			cmd = exec.Command("brew", "search", pkgs)
		default:
			if err := handler.OpError(macpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "windows":
		switch winpacman {
		case "choco", "winget", "scoop":
			cmd = exec.Command(winpacman, "search", pkgs)
		default:
			if err := handler.OpError(winpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "linux":
		switch pacman {
		case "apt", "dnf", "yum", "pkg", "zypper", "apk", "pkgin":
			cmd = exec.Command(pacman, "search", pkgs)
		case "apt-get":
			cmd = exec.Command("apt-cache", "search", pkgs)
		case "pacman":
			cmd = exec.Command(pacman, "-Ss", pkgs)
		case "xbps-install":
			cmd = exec.Command("xbps-query", "-Rs", pkgs)
		case "emerge":
			cmd = exec.Command("emerge", "--search", pkgs)
		case "pkg_add":
			cmd = exec.Command("pkg_info", "-Q", pkgs)
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
