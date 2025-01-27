package system

import (
	"fmt"
	"os"
	"path/filepath"
)

func RemovingPackageManager(pkg string) {
	fmt.Printf("Remove %s...\n", pkg)
}

func UninstallAIOPackageManager() {
	distro := Distro()
	cmdList := []string{"install", "reinstall", "update", "upgrade", "search", "updateupgrade", "uninstall", "detail", "hold", "unhold", "purge", "orphan", "list"}
	switch distro {
	case "termux":
		for _, cmd := range cmdList {
			binPath := filepath.Join("/data/data/com.termux/files/usr/bin", cmd)
			if PathExists(binPath) {
				if err := os.Remove(binPath); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				RemovingPackageManager(cmd)
			} else {
				fmt.Printf("%s is not installed\n", cmd)
				os.Exit(1)
			}
		}
	case "macos":
		for _, cmd := range cmdList {
			binPath := filepath.Join("/usr/local/bin", cmd)
			if PathExists(binPath) {
				if err := os.Remove(binPath); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				RemovingPackageManager(cmd)
			} else {
				fmt.Printf("%s is not installed\n", cmd)
				os.Exit(1)
			}
		}
	case "windows":
		for _, cmd := range cmdList {
			binPath := filepath.Join("C:\\Windows\\System32", cmd)
			if PathExists(binPath) {
				if err := os.Remove(binPath); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				RemovingPackageManager(cmd)
			} else {
				fmt.Printf("%s is not installed\n", cmd)
				os.Exit(1)
			}
		}
	case "linux":
		pathOne := filepath.Join("/usr/bin")
		pathTwo := filepath.Join("/bin")
		pathThree := filepath.Join("/sbin")
		pathFour := filepath.Join("/usr/sbin")
		pathFive := filepath.Join("/usr/local/bin")
		if PathExists(pathOne) {
			for _, cmd := range cmdList {
				binPath := filepath.Join(pathOne, cmd)
				if PathExists(binPath) {
					if err := os.Remove(binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					RemovingPackageManager(cmd)
				} else {
					fmt.Printf("%s is not installed\n", cmd)
					os.Exit(1)
				}
			}
		} else if PathExists(pathTwo) {
			for _, cmd := range cmdList {
				binPath := filepath.Join(pathTwo, cmd)
				if PathExists(binPath) {
					if err := os.Remove(binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					RemovingPackageManager(cmd)
				} else {
					fmt.Printf("%s is not installed\n", cmd)
					os.Exit(1)
				}
			}
		} else if PathExists(pathThree) {
			for _, cmd := range cmdList {
				binPath := filepath.Join(pathThree, cmd)
				if PathExists(binPath) {
					if err := os.Remove(binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					RemovingPackageManager(cmd)
				} else {
					fmt.Printf("%s is not installed\n", cmd)
					os.Exit(1)
				}
			}
		} else if PathExists(pathFour) {
			for _, cmd := range cmdList {
				binPath := filepath.Join(pathFour, cmd)
				if PathExists(binPath) {
					if err := os.Remove(binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					RemovingPackageManager(cmd)
				} else {
					fmt.Printf("%s is not installed\n", cmd)
					os.Exit(1)
				}
			}
		} else if PathExists(pathFive) {
			for _, cmd := range cmdList {
				binPath := filepath.Join(pathFive, cmd)
				if PathExists(binPath) {
					if err := os.Remove(binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					RemovingPackageManager(cmd)
				} else {
					fmt.Printf("%s is not installed\n", cmd)
					os.Exit(1)
				}
			}
		}
	default:
		fmt.Println("Error: Unknown distro")
		os.Exit(1)
	}
}
