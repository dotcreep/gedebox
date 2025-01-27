package pkgman_test

import (
	"testing"

	"github.com/dotcreep/gedebox/internal/cli/pkgman"
)

func TestDetail(t *testing.T) {
	packagename := "python"
	pkgman.Detail(packagename)
}
