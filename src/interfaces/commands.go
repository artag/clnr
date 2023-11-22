package interfaces

import (
	d "github.com/artag/clnr/domain"
)

type ISaveRuleCommand interface {
	Execute(request *SaveRuleRequest) error
}

type SaveRuleRequest struct {
	Filename string
	Rule     *d.Rule
}

type IClearCliCommand interface {
	Execute()
}

type ISelectOptionCliCommand interface {
	Execute(minOptionNumber, maxOptionNumber int, printCaption func()) (option int)
}

type IPrintCliCommand interface {
	Printf(format string, a ...any) (int, error)
	Println(str string) (int, error)
	PrintStrings(slice []string)
	PrintCaptionStrings(caption string, slice []string)
	PrintCaptionIndexStrings(caption string, slice []string)
}

type IExcludeDirsCliCommand interface {
	Execute(directories []string) []string
}

type IExcludeFilesCliCommand interface {
	Execute(files []string) []string
}

type ISelectDirsCliCommand interface {
	Execute(directories []string) (selectedDirectories []string, quit bool)
}

type ISelectFilesCliCommand interface {
	Execute(files []string) (selectedFiles []string, quit bool)
}

type IDeleteDirsCommand interface {
	Execute(dirs []string)
}

type IDeleteFilesCommand interface {
	Execute(files []string)
}
