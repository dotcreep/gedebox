package system

import (
	"os/user"
)

func User() string {
	u, _ := user.Current()
	if u.Uid == "0" {
		return ""
	} else {
		return "sudo"
	}
}
