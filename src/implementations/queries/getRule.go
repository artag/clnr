package implementations

import (
	"encoding/json"
	"io/fs"
	"os"

	u "github.com/artag/clnr/common"
	d "github.com/artag/clnr/domain"
	i "github.com/artag/clnr/interfaces"
)

type GetRule struct {
	_readFile   func(string) ([]byte, error)
	_fileExists func(string) (fs.FileInfo, error)
}

func NewGetRule() *GetRule {
	return NewGetRuleInternal(os.ReadFile, os.Stat)
}

func NewGetRuleInternal(
	readFile func(string) ([]byte, error),
	fileExists func(string) (fs.FileInfo, error),
) *GetRule {
	u.AssertNotNil(readFile)
	u.AssertNotNil(fileExists)
	return &GetRule{readFile, fileExists}
}

func (q *GetRule) Execute(request *i.GetRuleRequest) *i.GetRuleResponse {
	if !q.fileExists(request.RuleFilename) {
		return &i.GetRuleResponse{Rule: d.CreateEmptyRule(), Error: nil}
	}

	file, err := q._readFile(request.RuleFilename)
	if err != nil {
		return &i.GetRuleResponse{Rule: nil, Error: err}
	}

	if len(file) == 0 {
		return &i.GetRuleResponse{Rule: nil, Error: d.ErrEmptyRuleFile}
	}

	rule := d.Rule{}
	err = json.Unmarshal(file, &rule)
	if err != nil {
		return &i.GetRuleResponse{Rule: nil, Error: err}
	}
	return &i.GetRuleResponse{Rule: &rule, Error: nil}
}

func (q *GetRule) fileExists(filename string) bool {
	if _, err := q._fileExists(filename); err == nil {
		return true

	} else {
		return false
	}
}
