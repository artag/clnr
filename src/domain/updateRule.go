package domain

type UpdateRuleFunc func(*Args, *Rule) *Rule

type UpdateRule struct {
	updates []UpdateRuleFunc
}

func NewUpdateRule() *UpdateRule {
	return &UpdateRule{
		[]UpdateRuleFunc{
			addIncludeDirEqual,
			addIncludeDirContain,
			addIncludeDirStart,
			addIncludeDirEnd,

			addExcludeDirEqual,
			addExcludeDirContain,
			addExcludeDirStart,
			addExcludeDirEnd,

			addIncludeFileEqual,
			addIncludeFileContain,
			addIncludeFileStart,
			addIncludeFileEnd,

			addExcludeFileEqual,
			addExcludeFileContain,
			addExcludeFileStart,
			addExcludeFileEnd,

			addRootDirectory,
			addRootDirectories,
		}}
}

func (u *UpdateRule) Execute(args *Args, rule *Rule) *Rule {
	r := rule
	for _, update := range u.updates {
		r = update(args, r)
	}
	return r
}

func addIncludeDirEqual(args *Args, rule *Rule) *Rule {
	if args.IsFileRule || !args.Include || args.Equal == "" {
		return rule
	}
	rule.Directory.Include.Equal = append(rule.Directory.Include.Equal, args.Equal)
	return rule
}

func addIncludeDirContain(args *Args, rule *Rule) *Rule {
	if args.IsFileRule || !args.Include || args.Contains == "" {
		return rule
	}
	rule.Directory.Include.Contain = append(rule.Directory.Include.Contain, args.Contains)
	return rule
}

func addIncludeDirStart(args *Args, rule *Rule) *Rule {
	if args.IsFileRule || !args.Include || args.StartWith == "" {
		return rule
	}
	rule.Directory.Include.Start = append(rule.Directory.Include.Start, args.StartWith)
	return rule
}

func addIncludeDirEnd(args *Args, rule *Rule) *Rule {
	if args.IsFileRule || !args.Include || args.Contains == "" {
		return rule
	}
	rule.Directory.Include.End = append(rule.Directory.Include.End, args.EndWith)
	return rule
}

func addExcludeDirEqual(args *Args, rule *Rule) *Rule {
	if args.IsFileRule || args.Include || args.Equal == "" {
		return rule
	}
	rule.Directory.Exclude.Equal = append(rule.Directory.Exclude.Equal, args.Equal)
	return rule
}

func addExcludeDirContain(args *Args, rule *Rule) *Rule {
	if args.IsFileRule || args.Include || args.Contains == "" {
		return rule
	}
	rule.Directory.Exclude.Contain = append(rule.Directory.Exclude.Contain, args.Contains)
	return rule
}

func addExcludeDirStart(args *Args, rule *Rule) *Rule {
	if args.IsFileRule || args.Include || args.StartWith == "" {
		return rule
	}
	rule.Directory.Exclude.Start = append(rule.Directory.Exclude.Start, args.StartWith)
	return rule
}

func addExcludeDirEnd(args *Args, rule *Rule) *Rule {
	if args.IsFileRule || args.Include || args.EndWith == "" {
		return rule
	}
	rule.Directory.Exclude.End = append(rule.Directory.Exclude.End, args.EndWith)
	return rule
}

func addIncludeFileEqual(args *Args, rule *Rule) *Rule {
	if !args.IsFileRule || !args.Include || args.Equal == "" {
		return rule
	}
	rule.File.Include.Equal = append(rule.File.Include.Equal, args.Equal)
	return rule
}

func addIncludeFileContain(args *Args, rule *Rule) *Rule {
	if !args.IsFileRule || !args.Include || args.Contains == "" {
		return rule
	}
	rule.File.Include.Contain = append(rule.File.Include.Contain, args.Contains)
	return rule
}

func addIncludeFileStart(args *Args, rule *Rule) *Rule {
	if !args.IsFileRule || !args.Include || args.StartWith == "" {
		return rule
	}
	rule.File.Include.Start = append(rule.File.Include.Start, args.StartWith)
	return rule
}

func addIncludeFileEnd(args *Args, rule *Rule) *Rule {
	if !args.IsFileRule || !args.Include || args.EndWith == "" {
		return rule
	}
	rule.File.Include.End = append(rule.File.Include.End, args.EndWith)
	return rule
}

func addExcludeFileEqual(args *Args, rule *Rule) *Rule {
	if !args.IsFileRule || args.Include || args.Equal == "" {
		return rule
	}
	rule.File.Exclude.Equal = append(rule.File.Exclude.Equal, args.Equal)
	return rule
}

func addExcludeFileContain(args *Args, rule *Rule) *Rule {
	if !args.IsFileRule || args.Include || args.Contains == "" {
		return rule
	}
	rule.File.Exclude.Contain = append(rule.File.Exclude.Contain, args.Contains)
	return rule
}

func addExcludeFileStart(args *Args, rule *Rule) *Rule {
	if !args.IsFileRule || args.Include || args.StartWith == "" {
		return rule
	}
	rule.File.Exclude.Start = append(rule.File.Exclude.Start, args.StartWith)
	return rule
}

func addExcludeFileEnd(args *Args, rule *Rule) *Rule {
	if !args.IsFileRule || args.Include || args.EndWith == "" {
		return rule
	}
	rule.File.Exclude.End = append(rule.File.Exclude.End, args.EndWith)
	return rule
}

func addRootDirectory(args *Args, rule *Rule) *Rule {
	if args.RootDir == "" {
		return rule
	}
	rule.RootDirectories = append(rule.RootDirectories, args.RootDir)
	return rule
}

func addRootDirectories(args *Args, rule *Rule) *Rule {
	if len(args.RootDirectories) < 1 {
		return rule
	}
	rule.RootDirectories = append(rule.RootDirectories, args.RootDirectories...)
	return rule
}
