package pkgman

import (
	"gedebox/system"
	"os"
	"path/filepath"
	"strings"
)

func Main() {
	programName := filepath.Base(os.Args[0])
	pkg := os.Args[1:]
	pkgs := strings.Join(pkg, " ")

	switch programName {
	case "install", "install.exe":
		Install(pkgs)
	case "reinstall", "reinstall.exe":
		Reinstall(pkgs)
	case "update", "update.exe":
		Update()
	case "upgrade", "upgrade.exe":
		Upgrade()
	case "search", "search.exe":
		Search(pkgs)
	case "updateupgrade", "updateupgrade.exe":
		UpdateANDUpgrades()
	case "uninstall", "uninstall.exe":
		Uninstall(pkgs)
	case "detail", "detail.exe":
		Detail(pkgs)
	case "hold", "hold.exe":
		Hold(pkgs)
	case "unhold", "unhold.exe":
		Unhold(pkgs)
	case "purge", "purge.exe":
		Purge(pkgs)
	case "orphan", "orphan.exe":
		Orphan()
	case "list", "list.exe":
		List()
	case "auri":
		AUR("auri", pkgs)
	case "aurinc":
		AUR("aurinc", pkgs)
	case "auru":
		AUR("auru", pkgs)
	case "auruu":
		AUR("auruu", pkgs)
	case "aurs":
		AUR("aurs", pkgs)
	case "aurr":
		AUR("aurr", pkgs)
	case "snapi":
		SNAP("snapi", pkgs)
	case "snapu":
		SNAP("snapu", pkgs)
	case "snapv":
		SNAP("snapv", pkgs)
	case "snaps":
		SNAP("snaps", pkgs)
	case "snapl":
		SNAP("snapl", pkgs)
	case "snapla":
		SNAP("snapla", pkgs)
	case "snapon":
		SNAP("snapon", pkgs)
	case "snapoff":
		SNAP("snapoff", pkgs)
	case "snapr":
		SNAP("snapr", pkgs)
	default:
		system.Help()
	}
}
