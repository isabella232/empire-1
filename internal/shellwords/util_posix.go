// +build !windows

package shellwords

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func shellRun(line string) (string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "sh"
	}
	b, err := exec.Command(shell, "-c", line).Output()
	if err != nil {
		return "", errors.New(err.Error() + ":" + string(b))
	}
	return strings.TrimSpace(string(b)), nil
}
