package domain

type Rule struct {
	Directory struct {
		Include struct {
			Equal   []string `json:"equal"`
			Contain []string `json:"contain"`
			Start   []string `json:"startWith"`
			End     []string `json:"endWith"`
		} `json:"include"`
		Exclude struct {
			Equal   []string `json:"equal"`
			Contain []string `json:"contain"`
			Start   []string `json:"start"`
			End     []string `json:"end"`
		} `json:"exclude"`
	} `json:"directory"`
	File struct {
		Include struct {
			Equal   []string `json:"equal"`
			Contain []string `json:"contain"`
			Start   []string `json:"start"`
			End     []string `json:"end"`
		} `json:"include"`
		Exclude struct {
			Equal   []string `json:"equal"`
			Contain []string `json:"contain"`
			Start   []string `json:"start"`
			End     []string `json:"end"`
		} `json:"exclude"`
	} `json:"file"`
	RootDirectories []string `json:"rootDirectories"`
}

func CreateEmptyRule() *Rule {
	rule := Rule{}

	rule.Directory.Include.Equal = make([]string, 0)
	rule.Directory.Include.Contain = make([]string, 0)
	rule.Directory.Include.Start = make([]string, 0)
	rule.Directory.Include.End = make([]string, 0)

	rule.Directory.Exclude.Equal = make([]string, 0)
	rule.Directory.Exclude.Contain = make([]string, 0)
	rule.Directory.Exclude.Start = make([]string, 0)
	rule.Directory.Exclude.End = make([]string, 0)

	rule.File.Include.Equal = make([]string, 0)
	rule.File.Include.Contain = make([]string, 0)
	rule.File.Include.Start = make([]string, 0)
	rule.File.Include.End = make([]string, 0)

	rule.File.Exclude.Equal = make([]string, 0)
	rule.File.Exclude.Contain = make([]string, 0)
	rule.File.Exclude.Start = make([]string, 0)
	rule.File.Exclude.End = make([]string, 0)

	rule.RootDirectories = make([]string, 0)

	return &rule
}

func (r *Rule) HasDirectoryData() bool {
	if len(r.Directory.Include.Equal) > 0 ||
		len(r.Directory.Include.Contain) > 0 ||
		len(r.Directory.Include.Start) > 0 ||
		len(r.Directory.Include.End) > 0 {
		return true
	}

	return false
}

func (r *Rule) HasFileData() bool {
	if len(r.File.Include.Equal) > 0 ||
		len(r.File.Include.Contain) > 0 ||
		len(r.File.Include.Start) > 0 ||
		len(r.File.Include.End) > 0 {
		return true
	}

	return false
}

func (r *Rule) HasRootDirectories() bool {
	return len(r.RootDirectories) > 0
}

func (r *Rule) NotEmptyRootDirectories() error {
	if r.HasRootDirectories() {
		return nil
	}
	return ErrEmptyRootDirectories
}
