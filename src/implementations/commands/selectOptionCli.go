package implementations

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	u "github.com/artag/clnr/common"
	i "github.com/artag/clnr/interfaces"
)

type SelectOptionCli struct {
	_clearConsole i.IClearCliCommand
	_printer      i.IPrintCliCommand
	_reader       *bufio.Reader
}

func NewSelectOptionCli(
	clearConsole i.IClearCliCommand,
	printer i.IPrintCliCommand,
) *SelectOptionCli {
	return NewSelectOptionCliInternal(
		clearConsole,
		printer,
		bufio.NewReader(os.Stdin),
	)
}

func NewSelectOptionCliInternal(
	clearConsole i.IClearCliCommand,
	printer i.IPrintCliCommand,
	reader *bufio.Reader,
) *SelectOptionCli {
	u.AssertNotNil(clearConsole)
	u.AssertNotNil(printer)
	u.AssertNotNil(reader)

	return &SelectOptionCli{clearConsole, printer, reader}
}

func (c *SelectOptionCli) Execute(min, max int, printCaption func()) int {
	errMessage := ""
	for {
		c._clearConsole.Execute()
		printCaption()
		c._printer.Println("---")
		if errMessage != "" {
			c._printer.Printf("Error: %s\n", errMessage)
		}
		c._printer.Printf("Select an option: ")
		input, err := c._reader.ReadString('\n')
		c._printer.Printf("")
		if err != nil {
			errMessage = "Wrong input. Enter integer number."
			continue
		}

		optionStr := strings.TrimSuffix(input, "\n")
		option, err := strconv.Atoi(optionStr)
		if err != nil {
			errMessage = fmt.Sprintf("Wrong input '%s'. Enter integer number.", input)
			continue
		}

		if min <= option && option <= max {
			errMessage = ""
			return option
		}

		errMessage = fmt.Sprintf("Wrong selection '%d'. Enter integer number from %d to %d.",
			option, min, max)
	}
}
