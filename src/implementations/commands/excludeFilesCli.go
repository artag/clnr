package implementations

import (
	u "github.com/artag/clnr/common"
	i "github.com/artag/clnr/interfaces"
)

type ExcludeFilesCli struct {
	_clearConsole i.IClearCliCommand
	_printer      i.IPrintCliCommand
	_selectOption i.ISelectOptionCliCommand
}

func NewExcludeFilesCli(
	clearConsole i.IClearCliCommand,
	printer i.IPrintCliCommand,
	selectOption i.ISelectOptionCliCommand,
) *ExcludeFilesCli {
	u.AssertNotNil(clearConsole)
	u.AssertNotNil(printer)
	u.AssertNotNil(selectOption)

	return &ExcludeFilesCli{clearConsole, printer, selectOption}
}

func (c *ExcludeFilesCli) Execute(files []string) (remainingFiles []string) {
	selectedFiles := files
	for {
		maxRow := len(selectedFiles)
		c._clearConsole.Execute()

		opt := c._selectOption.Execute(
			0, maxRow,
			func() {
				c._printer.PrintCaptionIndexStrings("Files to delete:", selectedFiles)
				c._printer.Printf("Choose an option:\n"+
					"- Enter row number (from 1 to %d) to exclude file\n"+
					"- Enter 0 to return\n", maxRow)
			})

		if opt == 0 {
			return selectedFiles
		}
		if 0 < opt && opt <= maxRow {
			idx := opt - 1
			selectedFiles = u.Remove(selectedFiles, idx)
		}
	}
}
