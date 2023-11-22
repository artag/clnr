package implementations

import (
	"fmt"
	"os"

	u "github.com/artag/clnr/common"
	i "github.com/artag/clnr/interfaces"
)

type DeleteFiles struct {
	_deleteFile func(string) error
	_printer    i.IPrintCliCommand
}

func NewDeleteFiles(printer i.IPrintCliCommand) *DeleteFiles {
	return NewDeleteFilesInternal(os.Remove, printer)
}

func NewDeleteFilesInternal(
	deleteFile func(string) error,
	printer i.IPrintCliCommand,
) *DeleteFiles {
	u.AssertNotNil(deleteFile)
	u.AssertNotNil(printer)

	return &DeleteFiles{deleteFile, printer}
}

func (c *DeleteFiles) Execute(files []string) {
	if files == nil || len(files) < 1 {
		c._printer.Println("There is no files to delete")
		return
	}

	for _, file := range files {
		msg := fmt.Sprintf("Deleting '%s'. ", file)
		if err := c.delete(file); err != nil {
			msg = msg + err.Error()
		} else {
			msg = msg + "Success."
		}
		c._printer.Println(msg)
	}
}

func (c *DeleteFiles) delete(file string) error {
	if err := c._deleteFile(file); err != nil {
		return err
	}
	return nil
}
