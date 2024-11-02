package pkgman

import (
	"fmt"
	"gedebox/handler"
	"gedebox/system"
	"os"
	"os/exec"
)

func Reinstall(pkgs string) {
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
		cmd = exec.Command("pkg", "reinstall", "-y", pkgs)
	case "macos":
		switch macpacman {
		case "brew":
			cmd = exec.Command("brew", "reinstall", pkgs)
		case "port":
			Uninstall(pkgs)
			Install(pkgs)
		default:
			if err := handler.OpError(macpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "windows":
		switch winpacman {
		case "choco":
			cmd = exec.Command("choco", "install", pkgs)
		case "winget":
			cmd = exec.Command("winget", "install", pkgs)
		case "scoop":
			cmd = exec.Command("scoop", "install", pkgs)
		default:
			if err := handler.OpError(winpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "linux":
		switch pacman {
		case "apt", "apt-get", "dnf", "yum", "pkgin":
			if su == "" {
				cmd = exec.Command(pacman, "reinstall", "-y", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "reinstall", "-y", pkgs)
			}
		case "pkg":
			if su == "" {
				cmd = exec.Command(pacman, "-y", "-f", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "-y", "-f", pkgs)
			}
		case "zypper":
			if su == "" {
				cmd = exec.Command(pacman, "install", "--force", pkgs)
			} else {
				cmd = exec.Command(pacman, "install", "--force", pkgs)
			}
		case "pacman":
			if su == "" {
				cmd = exec.Command(pacman, "-S", "--force", "--noconfirm", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "-S", "--force", "--noconfirm", pkgs)
			}
		case "xbps-install":
			if su == "" {
				cmd = exec.Command(pacman, "-f", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "-f", pkgs)
			}
		case "apk":
			if su == "" {
				cmd = exec.Command(pacman, "fix", "--force", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "fix", "--force", pkgs)
			}
		case "emerge":
			Uninstall(pkgs)
			Install(pkgs)
		case "pkg_add":
			Uninstall(pkgs)
			Install(pkgs)
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
