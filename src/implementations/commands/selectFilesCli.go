package implementations

import (
	u "github.com/artag/clnr/common"
	i "github.com/artag/clnr/interfaces"
)

type SelectFilesCli struct {
	_clearConsole i.IClearCliCommand
	_excludeFiles i.IExcludeFilesCliCommand
	_printer      i.IPrintCliCommand
	_selectOption i.ISelectOptionCliCommand
}

func NewSelectFilesCli(
	clearConsole i.IClearCliCommand,
	excludeFiles i.IExcludeFilesCliCommand,
	printer i.IPrintCliCommand,
	selectOption i.ISelectOptionCliCommand,
) *SelectFilesCli {
	u.AssertNotNil(clearConsole)
	u.AssertNotNil(excludeFiles)
	u.AssertNotNil(printer)
	u.AssertNotNil(selectOption)

	return &SelectFilesCli{clearConsole, excludeFiles, printer, selectOption}
}

func (c *SelectFilesCli) Execute(files []string) (toDelete []string, quit bool) {
	if len(files) < 1 {
		return files, false
	}

	f := files
	for {
		c._clearConsole.Execute()

		opt := c._selectOption.Execute(1, 3,
			func() {
				c._printer.PrintCaptionIndexStrings("Files to delete:", f)
				c._printer.Println("---")
				c._printer.Println("Choose an option:\n" +
					"1. Delete all listed files\n" +
					"2. Exclude files from delete list\n" +
					"3. Quit")
			})
		switch opt {
		case 1:
			return f, false
		case 2:
			f = c._excludeFiles.Execute(files)
		case 3:
			return make([]string, 0), true
		default:
			panic("Invalid operation")
		}
	}
}
