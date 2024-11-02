package pkgman

import (
	"fmt"
	"gedebox/handler"
	"gedebox/system"
	"os"
	"os/exec"
)

func Detail(pkg string) {
	if pkg == "" {
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
		cmd = exec.Command("pkg", "info", pkg)
	case "macos":
		switch macpacman {
		case "brew":
			cmd = exec.Command("brew", "info", pkg)
		case "port":
			cmd = exec.Command("port", "info", pkg)
		default:
			if err := handler.OpError(macpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "windows":
		switch winpacman {
		case "choco":
			cmd = exec.Command("choco", "info", pkg)
		case "winget":
			cmd = exec.Command("winget", "show", pkg)
		case "scoop":
			cmd = exec.Command("scoop", "info", pkg)
		default:
			if err := handler.OpError(winpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "linux":
		switch pacman {
		case "apt":
			cmd = exec.Command(pacman, "show", pkg)
		case "apt-get":
			cmd = exec.Command("apt-cache", "show", pkg)
		case "pkg", "dnf", "yum", "zypper":
			cmd = exec.Command(pacman, "info", pkg)
		case "pacman":
			cmd = exec.Command(pacman, "-Qi", pkg)
		case "xbps-install":
			cmd = exec.Command("xbps-query", "-Ri", pkg)
		case "apk":
			cmd = exec.Command(pacman, "info", pkg)
		case "emerge":
			cmd = exec.Command(pacman, "--info", pkg)
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
