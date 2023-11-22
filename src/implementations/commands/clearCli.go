package implementations

import (
	"os"
	"os/exec"
	"runtime"

	u "github.com/artag/clnr/common"
)

var supportedOS = map[string]bool{
	"linux":   true,
	"windows": true,
}

type ClearCli struct {
	_goos  string
	_clear map[string]func()
}

func NewClearCli() (*ClearCli, error) {
	clear := make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	return NewClearCliInternal(runtime.GOOS, clear)
}

func NewClearCliInternal(goos string, clearFunc map[string]func()) (*ClearCli, error) {
	u.AssertNotEmptyString(goos)
	u.AssertNotNil(clearFunc)

	if err := isCurrentOSSupported(goos); err != nil {
		return nil, err
	}

	return &ClearCli{goos, clearFunc}, nil
}

func (c *ClearCli) Execute() {
	clearConsole := c._clear[c._goos]
	clearConsole()
}

func isCurrentOSSupported(goos string) error {
	_, yes := supportedOS[goos]
	if yes {
		return nil
	}
	return ErrOSNotSupported
}
