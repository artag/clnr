package common_test

import (
	"reflect"
	"testing"

	"github.com/artag/clnr/common"
)

func TestContainsPunctsOnly(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "string is empty",
			input:    "",
			expected: false,
		},
		{
			name:     "string is 'abc/'",
			input:    "abc/",
			expected: false,
		},
		{
			name:     "string is 'c:\\abc'",
			input:    "c:\\abc",
			expected: false,
		},
		{
			name:     "string is './'",
			input:    "./",
			expected: true,
		},
		{
			name:     "string is '.\\'",
			input:    ".\\",
			expected: true,
		},
		{
			name:     "string is '...'",
			input:    "...",
			expected: true,
		},
		{
			name:     "string is '/tmp'",
			input:    "/tmp",
			expected: false,
		},
		{
			name:     "string is './file'",
			input:    "./file",
			expected: false,
		},
		{
			name:     "string is '../tmp'",
			input:    "../tmp",
			expected: false,
		},
		{
			name:     "string is ' .. '",
			input:    " .. ",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				act := common.ContainsPunctsOnly(tc.input)

				if tc.expected != act {
					t.Errorf("got %v want %v", act, tc.expected)
				}
			})
	}
}

func TestGetSliceOrEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "input slice is nil",
			input:    nil,
			expected: make([]string, 0),
		},
		{
			name:     "input slice is empty",
			input:    make([]string, 0),
			expected: make([]string, 0),
		},
		{
			name:     "input slice has any items",
			input:    []string{"1", "2", "abc"},
			expected: []string{"1", "2", "abc"},
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				act := common.GetSliceOrEmpty[string](tc.input)
				assertSlices(t, tc.expected, act)
			})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []string
		idx      int
		expected []string
	}{
		{
			name:     "input slice is nil",
			slice:    nil,
			idx:      42,
			expected: make([]string, 0),
		},
		{
			name:     "input slice is empty",
			slice:    make([]string, 0),
			idx:      1,
			expected: make([]string, 0),
		},
		{
			name:     "input slice is empty index is zero",
			slice:    make([]string, 0),
			idx:      0,
			expected: make([]string, 0),
		},
		{
			name:     "index is negative",
			slice:    []string{"1", "2", "3"},
			idx:      -1,
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "index is out of range",
			slice:    []string{"1", "2", "3"},
			idx:      3,
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "remove first item",
			slice:    []string{"1", "2", "3"},
			idx:      0,
			expected: []string{"2", "3"},
		},
		{
			name:     "remove last item",
			slice:    []string{"1", "2", "3"},
			idx:      2,
			expected: []string{"1", "2"},
		},
		{
			name:     "remove middle item",
			slice:    []string{"1", "2", "3", "4", "5"},
			idx:      2,
			expected: []string{"1", "2", "4", "5"},
		},
		{
			name:     "remove single item",
			slice:    []string{"1"},
			idx:      0,
			expected: make([]string, 0),
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				act := common.Remove[string](tc.slice, tc.idx)
				assertSlices(t, tc.expected, act)
			})
	}
}

func TestSplitPath(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		expected []string
	}{
		{
			name:     "empty path",
			path:     "",
			expected: []string{},
		},
		{
			name:     "path whitespaces only",
			path:     " ",
			expected: []string{},
		},
		{
			name:     "path '.'",
			path:     ".",
			expected: []string{},
		},
		{
			name:     "path '\\'",
			path:     "\\",
			expected: []string{},
		},
		{
			name:     "path '.\\'",
			path:     ".\\",
			expected: []string{},
		},
		{
			name:     "path '..'",
			path:     "..",
			expected: []string{},
		},
		{
			name:     "path '..\\'",
			path:     "..\\",
			expected: []string{},
		},
		{
			name:     "path 'C:'",
			path:     "C:",
			expected: []string{},
		},
		{
			name:     "path 'C:\\'",
			path:     "C:\\",
			expected: []string{},
		},
		{
			name:     "path 'C:\\Home'",
			path:     "C:\\Home",
			expected: []string{"Home"},
		},
		{
			name:     "path 'D:\\Home\\'",
			path:     "D:\\Home\\",
			expected: []string{"Home"},
		},
		{
			name:     "path 'd:\\Home\\root dir'",
			path:     "d:\\Home\\root dir",
			expected: []string{"Home", "root dir"},
		},
		{
			name:     "path 'd:\\Home\\tst\\'",
			path:     "d:\\Home\\tst\\",
			expected: []string{"Home", "tst"},
		},
		{
			name:     "path 'd:\\Home\\tst dir\\file.log'",
			path:     "d:\\Home\\tst dir\\file.log",
			expected: []string{"Home", "tst dir", "file.log"},
		},
		{
			name:     "path 'd:\\Home\\tst\\file.log\\'",
			path:     "d:\\Home\\tst\\file.log\\",
			expected: []string{"Home", "tst", "file.log"},
		},
		{
			name:     "path '..\\files\\work\\'",
			path:     "..\\files\\work\\",
			expected: []string{"files", "work"},
		},
		{
			name:     "path '.\\some dir\\program.bat'",
			path:     ".\\some dir\\program.bat",
			expected: []string{"some dir", "program.bat"},
		},
		{
			name:     "path '..\\files\\work\\app.exe'",
			path:     "..\\files\\work\\app.exe",
			expected: []string{"files", "work", "app.exe"},
		},
		{
			name:     "path '/'",
			path:     "/",
			expected: []string{},
		},
		{
			name:     "path './'",
			path:     "./",
			expected: []string{},
		},
		{
			name:     "path '//'",
			path:     "//",
			expected: []string{},
		},
		{
			name:     "path '/tmp'",
			path:     "/tmp",
			expected: []string{"tmp"},
		},
		{
			name:     "path './tmp'",
			path:     "./tmp",
			expected: []string{"tmp"},
		},
		{
			name:     "path '../tmp'",
			path:     "../tmp",
			expected: []string{"tmp"},
		},
		{
			name:     "path './tmp/log'",
			path:     "./tmp/log",
			expected: []string{"tmp", "log"},
		},
		{
			name:     "path 'tmp/log/'",
			path:     "tmp/log/",
			expected: []string{"tmp", "log"},
		},
		{
			name:     "path 'tmp/log/soft.log'",
			path:     "tmp/log/soft.log",
			expected: []string{"tmp", "log", "soft.log"},
		},
		{
			name:     "path '../tmp/log/soft.log/'",
			path:     "../tmp/log/soft.log/",
			expected: []string{"tmp", "log", "soft.log"},
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				act := common.SplitPath(tc.path)
				assertSlices(t, tc.expected, act)
			})
	}
}

func assertSlices(t *testing.T, want, got []string) {
	if reflect.DeepEqual(want, got) {
		return
	}

	t.Errorf("got %v want %v", got, want)
}
