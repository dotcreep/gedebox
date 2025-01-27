package pkgman_test

import (
	"testing"

	"github.com/dotcreep/gedebox/internal/cli/pkgman"
)

func TestUninstall(t *testing.T) {
	packagename := "btop"
	pkgman.Uninstall(packagename)
}
