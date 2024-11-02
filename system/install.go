package system

import (
	"fmt"
	"os"
	"path/filepath"
)

func ConfigurePackageManager(pkg string) {
	fmt.Printf("Configure %s...\n", pkg)
}

func ReconfigurePackageManager(pkg string) {
	fmt.Printf("Reconfigure %s...\n", pkg)
}

func FileAbs(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return absPath
}

func InstallAIOPackageManager() {
	distro := Distro()
	cmdList := []string{"install", "reinstall", "update", "upgrade", "search", "updateupgrade", "uninstall", "detail", "hold", "unhold", "purge", "orphan", "list"}
	linuxCmdList := []string{"auri", "aurinc", "auru", "auruu", "aurs", "aurr", "snapi", "snapu", "snapv", "snaps", "snapl", "snapla", "snapon", "snapoff", "snapr"}
	allCmdList := append(cmdList, linuxCmdList...)
	switch distro {
	case "termux":
		for _, cmd := range allCmdList {
			binPath := filepath.Join("/data/data/com.termux/files/usr/bin", cmd)
			bakPath := filepath.Join("/data/data/com.termux/files/usr/bin", cmd+".bak")
			if PathExists(binPath) {
				if !PathExists(bakPath) {
					if err := os.Rename(binPath, bakPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				} else {
					if err := os.Remove(binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
				if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				ReconfigurePackageManager(cmd)
			} else {
				if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				ConfigurePackageManager(cmd)
			}
		}
	case "macos":
		for _, cmd := range cmdList {
			binPath := filepath.Join("/usr/local/bin", cmd)
			bakPath := filepath.Join("/usr/local/bin", cmd+".bak")
			if PathExists(binPath) {
				if !PathExists(bakPath) {
					if err := os.Rename(binPath, bakPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				} else {
					if err := os.Remove(binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
				if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				ReconfigurePackageManager(cmd)
			} else {
				if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				ConfigurePackageManager(cmd)
			}
		}
	case "windows":
		for _, cmd := range cmdList {
			binPath := filepath.Join("C:\\Windows\\System32", cmd+".exe")
			bakPath := filepath.Join("C:\\Windows\\System32", cmd+".exe.bak")
			if PathExists(binPath) {
				if !PathExists(bakPath) {
					if err := os.Rename(binPath, bakPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				} else {
					if err := os.Remove(binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
				if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				ReconfigurePackageManager(cmd)
			} else {
				if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				ConfigurePackageManager(cmd)
			}
		}
	case "linux":
		pathOne := filepath.Join("/usr/bin")
		pathTwo := filepath.Join("/bin")
		pathThree := filepath.Join("/sbin")
		pathFour := filepath.Join("/usr/sbin")
		pathFive := filepath.Join("/usr/local/bin")
		if PathExists(pathOne) {
			for _, cmd := range allCmdList {
				binPath := filepath.Join(pathOne, cmd)
				bakPath := filepath.Join(pathOne, fmt.Sprintf("%s.bak", cmd))
				if PathExists(binPath) {
					if !PathExists(bakPath) {
						if err := os.Rename(binPath, bakPath); err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					} else {
						if err := os.Remove(binPath); err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					}
					if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					ReconfigurePackageManager(cmd)
				} else {
					if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					ConfigurePackageManager(cmd)
				}
			}
		} else if PathExists(pathTwo) {
			for _, cmd := range allCmdList {
				binPath := filepath.Join(pathTwo, cmd)
				bakPath := filepath.Join(pathTwo, fmt.Sprintf("%s.bak", cmd))
				if PathExists(binPath) {
					if !PathExists(bakPath) {
						if err := os.Rename(binPath, bakPath); err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					} else {
						if err := os.Remove(binPath); err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					}
					if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					ReconfigurePackageManager(cmd)
				} else {
					if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					ConfigurePackageManager(cmd)
				}
			}
		} else if PathExists(pathThree) {
			for _, cmd := range allCmdList {
				binPath := filepath.Join(pathThree, cmd)
				bakPath := filepath.Join(pathThree, fmt.Sprintf("%s.bak", cmd))
				if PathExists(binPath) {
					if !PathExists(bakPath) {
						if err := os.Rename(binPath, bakPath); err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					} else {
						if err := os.Remove(binPath); err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					}
					if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					ReconfigurePackageManager(cmd)
				} else {
					if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					ConfigurePackageManager(cmd)
				}
			}
		} else if PathExists(pathFour) {
			for _, cmd := range allCmdList {
				binPath := filepath.Join(pathFour, cmd)
				bakPath := filepath.Join(pathFour, fmt.Sprintf("%s.bak", cmd))
				if PathExists(binPath) {
					if !PathExists(bakPath) {
						if err := os.Rename(binPath, bakPath); err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					} else {
						if err := os.Remove(binPath); err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					}
					ReconfigurePackageManager(cmd)
					if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				} else {
					ConfigurePackageManager(cmd)
					if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
			}
		} else if PathExists(pathFive) {
			for _, cmd := range allCmdList {
				binPath := filepath.Join(pathFive, cmd)
				bakPath := filepath.Join(pathFive, fmt.Sprintf("%s.bak", cmd))
				if PathExists(binPath) {
					if !PathExists(bakPath) {
						if err := os.Rename(binPath, bakPath); err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					} else {
						if err := os.Remove(binPath); err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					}
					ReconfigurePackageManager(cmd)
					if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				} else {
					ConfigurePackageManager(cmd)
					if err := os.Symlink(FileAbs(os.Args[0]), binPath); err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
			}
		} else {
			fmt.Println("Unsupported distribution")
			os.Exit(1)
		}
	default:
		fmt.Println("Unsupported distribution")
		os.Exit(1)
	}
}
