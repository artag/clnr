package domain

import (
	utils "github.com/artag/clnr/common"
)

type Args struct {
	RuleFilename    string
	IsFileRule      bool
	Include         bool
	Equal           string
	Contains        string
	StartWith       string
	EndWith         string
	RootDir         string
	RootDirectories []string
}

func NewArgs(
	ruleFilename string,
	isFileRule bool,
	include bool,
	equal string,
	contain string,
	start string,
	end string,
	rootDir string,
	rootDirectories []string) *Args {
	return &Args{
		utils.GetValueOrEmpty(ruleFilename),
		isFileRule,
		include,
		utils.GetValueOrEmpty(equal),
		contain,
		start,
		end,
		utils.GetValueOrEmpty(rootDir),
		utils.GetSliceOrEmpty(rootDirectories),
	}
}

func (a *Args) RuleFilenameExists() error {
	var fn = a.RuleFilename
	if fn == "" || utils.ContainsWhitespacesOnly(fn) {
		return ErrEmptyRuleFilename
	}

	return nil
}

func (args *Args) SelectRuleAction() RuleAction {
	if args.hasSomeData() {
		return Save
	}
	return Execute
}

func (a *Args) hasSomeData() bool {
	if a.Equal != "" ||
		a.Contains != "" ||
		a.StartWith != "" ||
		a.EndWith != "" ||
		a.RootDir != "" {
		return true
	}

	return false
}
