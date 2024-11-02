package pkgman

import (
	"fmt"
	"gedebox/handler"
	"gedebox/system"
	"os"
	"os/exec"
)

func Purge(pkgs string) {
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
		cmd = exec.Command("pkg", "remove", pkgs)
	case "macos":
		switch macpacman {
		case "brew":
			cmd = exec.Command("brew", "uninstall", pkgs)
		case "port":
			cmd = exec.Command("port", "uninstall", pkgs)
		default:
			if err := handler.OpError(macpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "windows":
		switch winpacman {
		case "choco":
			cmd = exec.Command("choco", "uninstall", pkgs)
		case "winget":
			cmd = exec.Command("winget", "uninstall", pkgs)
		case "scoop":
			cmd = exec.Command("scoop", "uninstall", pkgs)
		default:
			if err := handler.OpError(winpacman); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "linux":
		switch pacman {
		case "apt", "apt-get":
			if su == "" {
				cmd = exec.Command(pacman, "purge", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "purge", pkgs)
			}
		case "dnf", "yum":
			if su == "" {
				cmd = exec.Command(pacman, "remove", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "remove", pkgs)
			}
		case "zypper":
			if su == "" {
				cmd = exec.Command(pacman, "remove", "--clean-deps", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "remove", "--clean-deps", pkgs)
			}
		case "pkg":
			if su == "" {
				cmd = exec.Command(pacman, "delete", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "delete", pkgs)
			}
		case "pkgin":
			if su == "" {
				cmd = exec.Command(pacman, "remove", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "remove", pkgs)
			}
		case "pkg_add":
			if su == "" {
				cmd = exec.Command("pkg_delete", pkgs)
			} else {
				cmd = exec.Command(su, "pkg_delete", pkgs)
			}
		case "pacman":
			if su == "" {
				cmd = exec.Command(pacman, "-Rns", "--noconfirm", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "-Rns", "--noconfirm", pkgs)
			}
		case "xbps-install":
			if su == "" {
				cmd = exec.Command("xbps-remove", "-Ry", pkgs)
			} else {
				cmd = exec.Command(su, "xbps-remove", "-Ry", pkgs)
			}
		case "apk":
			if su == "" {
				cmd = exec.Command(pacman, "del", "--purge", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "del", "--purge", pkgs)
			}
		case "emerge":
			if su == "" {
				cmd = exec.Command(pacman, "--depclean", "--verbose", pkgs)
			} else {
				cmd = exec.Command(su, pacman, "--depclean", "--verbose", pkgs)
			}
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
