package interfaces

import (
	d "github.com/artag/clnr/domain"
)

type IGetRuleQuery interface {
	Execute(request *GetRuleRequest) *GetRuleResponse
}

type GetRuleRequest struct {
	RuleFilename string
}
type GetRuleResponse struct {
	Rule  *d.Rule
	Error error
}

type ISearchDirsToDeleteQuery interface {
	Execute(rule *d.Rule) (dirsToDelete []string)
}

type ISearchFilesToDeleteQuery interface {
	Execute(rule *d.Rule) (filesToDelete []string)
}
