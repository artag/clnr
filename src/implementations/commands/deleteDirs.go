package implementations

import (
	"fmt"
	"os"

	u "github.com/artag/clnr/common"
	i "github.com/artag/clnr/interfaces"
)

type DeleteDirs struct {
	_deleteDir func(string) error
	_printer   i.IPrintCliCommand
}

func NewDeleteDirs(printer i.IPrintCliCommand) *DeleteDirs {
	return NewDeleteDirsInternal(os.RemoveAll, printer)
}

func NewDeleteDirsInternal(
	deleteDir func(string) error,
	printer i.IPrintCliCommand,
) *DeleteDirs {
	u.AssertNotNil(deleteDir)
	u.AssertNotNil(printer)

	return &DeleteDirs{deleteDir, printer}
}

func (c *DeleteDirs) Execute(dirs []string) {
	if dirs == nil || len(dirs) < 1 {
		c._printer.Println("There is no directories to delete")
		return
	}

	for _, dir := range dirs {
		msg := fmt.Sprintf("Deleting '%s'. ", dir)
		if err := c.delete(dir); err != nil {
			msg = msg + err.Error()
		} else {
			msg = msg + "Success."
		}
		c._printer.Println(msg)
	}
}

func (c *DeleteDirs) delete(dir string) error {
	if err := c._deleteDir(dir); err != nil {
		return err
	}
	return nil
}
