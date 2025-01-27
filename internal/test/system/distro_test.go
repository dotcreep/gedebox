package system_test

import (
	"fmt"
	"testing"

	"github.com/dotcreep/gedebox/internal/cli/system"
)

func TestDistro(t *testing.T) {
	distro := system.Distro()
	fmt.Println(distro)
}
