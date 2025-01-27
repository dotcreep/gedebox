package pkgman_test

import (
	"testing"

	"github.com/dotcreep/gedebox/internal/cli/pkgman"
)

func TestInstall(t *testing.T) {
	command := "btop"
	pkgman.Install(command)
}
