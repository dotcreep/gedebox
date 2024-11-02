package pkgman

import (
	"fmt"
	"gedebox/handler"
	"gedebox/system"
	"os"
	"os/exec"
)

func List() {
	distro := system.Distro()
	distro_id := system.DistroID()
	pacman := system.GetLinuxPackageManager(distro_id)
	winpacman := system.GetWindowsPackageManager()
	macpacman := system.GetMacPackageManager()
	var cmd *exec.Cmd
	switch distro {
	case "termux":
		cmd = exec.Command("pkg", "list")
	case "macos":
		switch macpacman {
		case "brew":
			cmd = exec.Command(macpacman, "list")
		case "port":
			cmd = exec.Command(macpacman, "installed")
		default:
			if err := handler.OpError(macpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Println("Package manager is not installed")
	case "windows":
		switch winpacman {
		case "choco":
			cmd = exec.Command("choco", "list", "--local-only")
		case "winget":
			cmd = exec.Command("winget", "list")
		case "scoop":
			cmd = exec.Command("scoop", "list")
		default:
			if err := handler.OpError(winpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "linux":
		switch pacman {
		case "apt", "apt-get":
			cmd = exec.Command(pacman, "list", "--installed")
		case "yum", "dnf":
			cmd = exec.Command(pacman, "list", "installed")
		case "pkg":
			cmd = exec.Command(pacman, "info")
		case "zypper":
			cmd = exec.Command(pacman, "search", "--installed-only")
		case "pacman":
			cmd = exec.Command(pacman, "-Q")
		case "xbps-install":
			cmd = exec.Command("xbps-query", "-l")
		case "apk":
			cmd = exec.Command(pacman, "list")
		case "emerge":
			cmd = exec.Command("emerge", "--list")
		case "pkg_add":
			cmd = exec.Command("pkg_info")
		case "pkgin":
			cmd = exec.Command("pkgin", "list")
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