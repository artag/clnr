package domain

import "errors"

var (
	ErrEmptyRuleFilename    = errors.New("empty rule filename")
	ErrMissingParameters    = errors.New("missing parameters")
	ErrUnknownAction        = errors.New("unknown action")
	ErrEmptyRuleFile        = errors.New("empty rule file")
	ErrEmptyRootDirectories = errors.New("empty root directories")
)
