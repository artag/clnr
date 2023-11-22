package implementations

import (
	u "github.com/artag/clnr/common"
	i "github.com/artag/clnr/interfaces"
)

type SelectDirsCli struct {
	_clearConsole i.IClearCliCommand
	_excludeDirs  i.IExcludeDirsCliCommand
	_printer      i.IPrintCliCommand
	_selectOption i.ISelectOptionCliCommand
}

func NewSelectDirsCli(
	clearConsole i.IClearCliCommand,
	excludeDirs i.IExcludeDirsCliCommand,
	printer i.IPrintCliCommand,
	selectOption i.ISelectOptionCliCommand,
) *SelectDirsCli {
	u.AssertNotNil(clearConsole)
	u.AssertNotNil(excludeDirs)
	u.AssertNotNil(printer)
	u.AssertNotNil(selectOption)

	return &SelectDirsCli{clearConsole, excludeDirs, printer, selectOption}
}

func (c *SelectDirsCli) Execute(directories []string) (selectedDirectories []string, quit bool) {
	if len(directories) < 1 {
		return directories, false
	}

	dirs := directories
	for {
		c._clearConsole.Execute()

		opt := c._selectOption.Execute(1, 3,
			func() {
				c._printer.PrintCaptionIndexStrings("Directories to delete:", dirs)
				c._printer.Println("---")
				c._printer.Println("Choose an option:\n" +
					"1. Delete all listed directories\n" +
					"2. Exclude directories from delete list\n" +
					"3. Quit")
			})
		switch opt {
		case 1:
			return dirs, false
		case 2:
			dirs = c._excludeDirs.Execute(dirs)
		case 3:
			return make([]string, 0), true
		default:
			panic("Invalid operation")
		}
	}
}
