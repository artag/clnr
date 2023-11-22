package implementations

import (
	"fmt"

	u "github.com/artag/clnr/common"
)

type PrintCli struct {
	_printf func(string, ...any) (int, error)
}

func NewPrintCli() *PrintCli {
	return NewPrintCliInternal(fmt.Printf)
}

func NewPrintCliInternal(printf func(string, ...any) (int, error)) *PrintCli {
	u.AssertNotNil(printf)
	return &PrintCli{printf}
}

func (c *PrintCli) Printf(format string, a ...any) (int, error) {
	return c._printf(format, a...)
}

func (c *PrintCli) Println(str string) (int, error) {
	return c.Printf("%s\n", str)
}

func (c *PrintCli) PrintStrings(slice []string) {
	if slice == nil {
		return
	}

	for _, item := range slice {
		c.Println(item)
	}
}

func (c *PrintCli) PrintCaptionStrings(caption string, slice []string) {
	c.Println(caption)
	if slice == nil {
		return
	}

	for _, item := range slice {
		c.Println(item)
	}
}

func (c *PrintCli) PrintCaptionIndexStrings(caption string, slice []string) {
	c.Println(caption)
	if slice == nil {
		fmt.Println("Empty")
		return
	}

	for i, item := range slice {
		num := i + 1
		c._printf("%d. %s\n", num, item)
	}
}
