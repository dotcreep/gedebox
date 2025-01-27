package pkgman_test

import (
	"testing"

	"github.com/dotcreep/gedebox/internal/cli/pkgman"
)

func TestUnhold(t *testing.T) {
	packagename := "btop"
	pkgman.Unhold(packagename)
}
