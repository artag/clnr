package implementations

import (
	"encoding/json"
	"io/fs"
	"os"

	u "github.com/artag/clnr/common"
	i "github.com/artag/clnr/interfaces"
)

type SaveRule struct {
	_writeFile func(string, []byte, fs.FileMode) error
}

func NewSaveRule() *SaveRule {
	return NewSaveRuleInternal(os.WriteFile)
}

func NewSaveRuleInternal(writeFileFunc func(string, []byte, fs.FileMode) error) *SaveRule {
	u.AssertNotNil(writeFileFunc)
	return &SaveRule{writeFileFunc}
}

func (c *SaveRule) Execute(request *i.SaveRuleRequest) error {
	js, err := json.Marshal(request.Rule)
	if err != nil {
		return err
	}

	return c._writeFile(request.Filename, js, 0644)
}
