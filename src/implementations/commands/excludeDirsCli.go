package implementations

import (
	u "github.com/artag/clnr/common"
	i "github.com/artag/clnr/interfaces"
)

type ExcludeDirsCli struct {
	_clearConsole i.IClearCliCommand
	_printer      i.IPrintCliCommand
	_selectOption i.ISelectOptionCliCommand
}

func NewExcludeDirsCli(
	clearConsole i.IClearCliCommand,
	printer i.IPrintCliCommand,
	selectOption i.ISelectOptionCliCommand,
) *ExcludeDirsCli {
	u.AssertNotNil(clearConsole)
	u.AssertNotNil(printer)
	u.AssertNotNil(selectOption)

	return &ExcludeDirsCli{clearConsole, printer, selectOption}
}

func (c *ExcludeDirsCli) Execute(directories []string) (remainingDirectories []string) {
	selectedDirs := directories
	for {
		maxRow := len(selectedDirs)
		c._clearConsole.Execute()

		opt := c._selectOption.Execute(
			0, maxRow,
			func() {
				c._printer.PrintCaptionIndexStrings("Directories to delete:", selectedDirs)
				c._printer.Printf("Choose an option:\n"+
					"- Enter row number (from 1 to %d) to exclude directory\n"+
					"- Enter 0 to return\n", maxRow)
			})

		if opt == 0 {
			return selectedDirs
		}
		if 0 < opt && opt <= maxRow {
			idx := opt - 1
			selectedDirs = u.Remove(selectedDirs, idx)
		}
	}
}
