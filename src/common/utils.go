package common

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode"
)

func AssertNotNil(a any) {
	if a != nil {
		return
	}

	msg := fmt.Sprintf("%s is nil", reflect.TypeOf(a))
	panic(msg)
}

func AssertNotEmptyString(str string) {
	if str == "" || ContainsWhitespacesOnly(str) {
		panic("input string is empty or contains whitespaces only")
	}
}

func GetValueOrEmpty(str string) string {
	if str == "" || ContainsWhitespacesOnly(str) {
		return ""
	}
	return str
}

func ContainsWhitespacesOnly(str string) bool {
	for _, c := range str {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// True - string contains punctuations characters only.
// Otherwise - false.
func ContainsPunctsOnly(str string) bool {
	if str == "" {
		return false
	}

	for _, c := range str {
		if !unicode.IsPunct(c) {
			return false
		}
	}
	return true
}

// Get slice if input slice is not empty.
// Get empty slice if input slice is nil.
func GetSliceOrEmpty[T any](slice []T) []T {
	if slice == nil || len(slice) < 1 {
		return make([]T, 0)
	}
	return slice
}

// Remove item from slice by index.
func Remove[T any](slice []T, idx int) []T {
	if slice == nil || len(slice) < 1 {
		return make([]T, 0)
	}

	if idx < 0 || len(slice) <= idx {
		return slice
	}

	return append(slice[:idx], slice[idx+1:]...)
}

// Split path by separator.
func SplitPath(path string) []string {
	var separator = string(os.PathSeparator)
	var replacer = strings.NewReplacer("\\", separator, "/", separator)

	path1 := replacer.Replace(path)
	split := strings.Split(path1, separator)
	result := make([]string, 0)
	for _, s := range split {
		if s == "" ||
			ContainsWhitespacesOnly(s) ||
			ContainsPunctsOnly(s) ||
			strings.Contains(s, ":") {
			continue
		}

		result = append(result, s)
	}

	return result
}
