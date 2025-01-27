package pkgman_test

import (
	"testing"

	"github.com/dotcreep/gedebox/internal/cli/pkgman"
)

func TestSearch(t *testing.T) {
	command := "python3"
	pkgman.Search(command)
}
