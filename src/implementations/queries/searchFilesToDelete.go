package implementations

import (
	"os"
	"path/filepath"
	"strings"

	u "github.com/artag/clnr/common"
	d "github.com/artag/clnr/domain"
)

type SearchFilesToDelete struct {
	_filepathwalk func(root string, fn filepath.WalkFunc) error
}

func NewSearchFilesToDelete() *SearchFilesToDelete {
	return NewSearchFilesToDeleteInternal(filepath.Walk)
}

func NewSearchFilesToDeleteInternal(
	fn func(root string, fn filepath.WalkFunc) error) *SearchFilesToDelete {
	u.AssertNotNil(fn)
	return &SearchFilesToDelete{fn}
}

func (q *SearchFilesToDelete) Execute(rule *d.Rule) []string {
	if !rule.HasFileData() || !rule.HasRootDirectories() {
		return make([]string, 0)
	}

	files := make([]string, 0)
	roots := rule.RootDirectories
	for _, root := range roots {
		q._filepathwalk(
			root,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if info.IsDir() {
					return nil
				}

				if !fileInInclude(rule, path) {
					return nil
				}

				if fileInExclude(rule, path) {
					return nil
				}

				files = append(files, path)
				return nil
			})
	}

	return files
}

func fileInInclude(rule *d.Rule, path string) bool {
	file := strings.ToLower(filepath.Base(path))

	for _, str := range rule.File.Include.Equal {
		if strings.ToLower(str) == file {
			return true
		}
	}

	for _, substr := range rule.File.Include.Contain {
		if strings.Contains(file, strings.ToLower(substr)) {
			return true
		}
	}

	for _, prefix := range rule.File.Include.Start {
		if strings.HasPrefix(file, strings.ToLower(prefix)) {
			return true
		}
	}

	for _, suffix := range rule.File.Include.End {
		if strings.HasSuffix(file, strings.ToLower(suffix)) {
			return true
		}
	}

	return false
}

func fileInExclude(rule *d.Rule, path string) bool {
	pathLower := strings.ToLower(path)
	split := u.SplitPath(pathLower)

	for _, file := range split {
		for _, str := range rule.File.Exclude.Equal {
			if strings.ToLower(str) == file {
				return true
			}
		}

		for _, substr := range rule.File.Exclude.Contain {
			if strings.Contains(file, strings.ToLower(substr)) {
				return true
			}
		}

		for _, prefix := range rule.File.Exclude.Start {
			if strings.HasPrefix(file, strings.ToLower(prefix)) {
				return true
			}
		}

		for _, suffix := range rule.File.Exclude.End {
			if strings.HasSuffix(file, strings.ToLower(suffix)) {
				return true
			}
		}
	}

	return false
}
