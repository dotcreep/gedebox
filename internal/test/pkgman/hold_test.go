package pkgman_test

import (
	"testing"

	"github.com/dotcreep/gedebox/internal/cli/pkgman"
)

func TestHold(t *testing.T) {
	packagename := "btop"
	pkgman.Hold(packagename)
}
