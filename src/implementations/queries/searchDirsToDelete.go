package implementations

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	u "github.com/artag/clnr/common"
	d "github.com/artag/clnr/domain"
)

type SearchDirsToDelete struct {
	_filepathwalk func(root string, fn filepath.WalkFunc) error
}

func NewSearchDirsToDelete() *SearchDirsToDelete {
	return NewSearchDirsToDeleteInternal(filepath.Walk)
}

func NewSearchDirsToDeleteInternal(
	fn func(root string, fn filepath.WalkFunc) error) *SearchDirsToDelete {
	u.AssertNotNil(fn)
	return &SearchDirsToDelete{fn}
}

func (q *SearchDirsToDelete) Execute(rule *d.Rule) []string {
	if !rule.HasDirectoryData() || !rule.HasRootDirectories() {
		return make([]string, 0)
	}

	dirs := make([]string, 0)
	roots := rule.RootDirectories
	for _, root := range roots {
		q._filepathwalk(
			root,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if !info.IsDir() {
					return nil
				}

				if !dirInInclude(rule, path) {
					return nil
				}

				if dirInExclude(rule, path) {
					return nil
				}

				dirs = append(dirs, path)
				return fs.SkipDir
			})
	}

	return dirs
}

func dirInInclude(rule *d.Rule, path string) bool {
	dir := strings.ToLower(filepath.Base(path))

	for _, str := range rule.Directory.Include.Equal {
		if strings.ToLower(str) == dir {
			return true
		}
	}

	for _, substr := range rule.Directory.Include.Contain {
		if strings.Contains(dir, strings.ToLower(substr)) {
			return true
		}
	}

	for _, prefix := range rule.Directory.Include.Start {
		if strings.HasPrefix(dir, strings.ToLower(prefix)) {
			return true
		}
	}

	for _, suffix := range rule.Directory.Include.End {
		if strings.HasSuffix(dir, strings.ToLower(suffix)) {
			return true
		}
	}

	return false
}

func dirInExclude(rule *d.Rule, path string) bool {
	pathLower := strings.ToLower(path)
	split := u.SplitPath(pathLower)

	for _, dir := range split {
		for _, str := range rule.Directory.Exclude.Equal {
			if strings.ToLower(str) == dir {
				return true
			}
		}

		for _, substr := range rule.Directory.Exclude.Contain {
			if strings.Contains(dir, strings.ToLower(substr)) {
				return true
			}
		}

		for _, prefix := range rule.Directory.Exclude.Start {
			if strings.HasPrefix(dir, strings.ToLower(prefix)) {
				return true
			}
		}

		for _, suffix := range rule.Directory.Exclude.End {
			if strings.HasSuffix(dir, strings.ToLower(suffix)) {
				return true
			}
		}
	}

	return false
}
