package system

import (
	"os/exec"
)

func CheckPackageManager(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func GetWindowsPackageManager() string {
	if CheckPackageManager("choco") {
		return "choco"
	} else if CheckPackageManager("scoop") {
		return "scoop"
	} else if CheckPackageManager("winget") {
		return "winget"
	}
	return ""
}

func GetMacPackageManager() string {
	if CheckPackageManager("brew") {
		return "brew"
	} else if CheckPackageManager("port") {
		return "port"
	}
	return ""
}

func GetLinuxPackageManager(distroID string) string {
	switch distroID {
	case "alpine":
		return "apk"
	case "arch":
		return "pacman"
	case "manjaro":
		return "pacman"
	case "gentoo":
		return "emerge"
	case "freebsd":
		return "pkg"
	case "openbsd":
		return "pkg_add"
	case "netbsd":
		return "pkgin"
	case "debian":
		return "apt"
	case "ubuntu":
		return "apt"
	case "fedora":
		return "dnf"
	case "centos":
		return "yum"
	case "alma-linux":
		return "yum"
	case "rocky-linux":
		return "dnf"
	case "oracle-linux":
		return "yum"
	case "rocky":
		return "dnf"
	case "rhel":
		return "yum"
	case "suse":
		return "zypper"
	case "opensuse-leap":
		return "zypper"
	case "sles":
		return "zypper"
	default:
		for _, pkg := range []string{"pacman", "apk", "zypper", "xbps-install", "pkg", "yum", "dnf", "apt", "apt-get"} {
			if Where(pkg) != "" {
				return pkg
			}
		}
	}
	return ""
}

func PackageManager() string {
	switch distro := Distro(); distro {
	case "termux":
		return "pkg"
	case "linux":
		distroID := DistroID()
		return GetLinuxPackageManager(distroID)
	case "macos":
		return "brew"
	case "windows":
		for _, pkg := range []string{"choco", "winget", "scoop"} {
			if Where(pkg) != "" {
				return pkg
			}
		}
	default:
		return "unknown"
	}
	return "unknown"
}
