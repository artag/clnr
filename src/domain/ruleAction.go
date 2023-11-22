package domain

type RuleAction int

const (
	Save    RuleAction = iota
	Execute RuleAction = iota
)
