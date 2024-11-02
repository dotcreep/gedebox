package system

import (
	"runtime"
)

func Arch() string {
	return runtime.GOARCH
}
