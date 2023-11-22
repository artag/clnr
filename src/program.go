package main

import (
	d "github.com/artag/clnr/domain"
	i "github.com/artag/clnr/interfaces"
)

type Program struct {
	_getRule             i.IGetRuleQuery
	_saveRule            i.ISaveRuleCommand
	_searchDirsToDelete  i.ISearchDirsToDeleteQuery
	_searchFilesToDelete i.ISearchFilesToDeleteQuery
	_selectDirectories   i.ISelectDirsCliCommand
	_selectFiles         i.ISelectFilesCliCommand
	_deleteDirs          i.IDeleteDirsCommand
	_deleteFiles         i.IDeleteFilesCommand
}

func (p *Program) Run(args *d.Args) error {
	if err := args.RuleFilenameExists(); err != nil {
		return err
	}

	getRule := p._getRule.Execute(&i.GetRuleRequest{RuleFilename: args.RuleFilename})
	if getRule.Error != nil {
		return getRule.Error
	}
	rule := getRule.Rule

	updateRule := d.NewUpdateRule()
	rule = updateRule.Execute(args, rule)
	action := args.SelectRuleAction()

	return p.executeAction(action, args, rule)
}

func (p *Program) executeAction(action d.RuleAction, args *d.Args, rule *d.Rule) error {
	switch action {
	case d.Save:
		if err := p._saveRule.Execute(&i.SaveRuleRequest{Filename: args.RuleFilename, Rule: rule}); err != nil {
			return err
		}
	case d.Execute:
		if err := rule.NotEmptyRootDirectories(); err != nil {
			return err
		}

		dirsToDelete := p._searchDirsToDelete.Execute(rule)
		filesToDelete := p._searchFilesToDelete.Execute(rule)

		selectedDirectories, quit := p._selectDirectories.Execute(dirsToDelete)
		if quit {
			return nil
		}

		selectedFiles, quit := p._selectFiles.Execute(filesToDelete)
		if quit {
			return nil
		}

		p._deleteDirs.Execute(selectedDirectories)
		p._deleteFiles.Execute(selectedFiles)
	default:
		return d.ErrUnknownAction
	}

	return nil
}
