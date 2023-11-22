package common_test

import (
	"reflect"
	"testing"

	"github.com/artag/clnr/common"
)

func TestSplitPath(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		expected []string
	}{
		{
			name:     "Empty path",
			path:     "",
			expected: []string{},
		},
		{
			name:     "Path whitespaces only",
			path:     " ",
			expected: []string{},
		},
		{
			name:     "Path '.'",
			path:     ".",
			expected: []string{},
		},
		{
			name:     "Path '\\'",
			path:     "\\",
			expected: []string{},
		},
		{
			name:     "Path '.\\'",
			path:     ".\\",
			expected: []string{},
		},
		{
			name:     "Path '..'",
			path:     "..",
			expected: []string{},
		},
		{
			name:     "Path '..\\'",
			path:     "..\\",
			expected: []string{},
		},
		{
			name:     "Path 'C:'",
			path:     "C:",
			expected: []string{},
		},
		{
			name:     "Path 'C:\\'",
			path:     "C:\\",
			expected: []string{},
		},
		{
			name:     "Path 'C:\\Home'",
			path:     "C:\\Home",
			expected: []string{"Home"},
		},
		{
			name:     "Path 'D:\\Home\\'",
			path:     "D:\\Home\\",
			expected: []string{"Home"},
		},
		{
			name:     "Path 'd:\\Home\\root dir'",
			path:     "d:\\Home\\root dir",
			expected: []string{"Home", "root dir"},
		},
		{
			name:     "Path 'd:\\Home\\tst\\'",
			path:     "d:\\Home\\tst\\",
			expected: []string{"Home", "tst"},
		},
		{
			name:     "Path 'd:\\Home\\tst dir\\file.log'",
			path:     "d:\\Home\\tst dir\\file.log",
			expected: []string{"Home", "tst dir", "file.log"},
		},
		{
			name:     "Path 'd:\\Home\\tst\\file.log\\'",
			path:     "d:\\Home\\tst\\file.log\\",
			expected: []string{"Home", "tst", "file.log"},
		},
		{
			name:     "Path '..\\files\\work\\'",
			path:     "..\\files\\work\\",
			expected: []string{"files", "work"},
		},
		{
			name:     "Path '.\\some dir\\program.bat'",
			path:     ".\\some dir\\program.bat",
			expected: []string{"some dir", "program.bat"},
		},
		{
			name:     "Path '..\\files\\work\\app.exe'",
			path:     "..\\files\\work\\app.exe",
			expected: []string{"files", "work", "app.exe"},
		},
		{
			name:     "Path '/'",
			path:     "/",
			expected: []string{},
		},
		{
			name:     "Path './'",
			path:     "./",
			expected: []string{},
		},
		{
			name:     "Path '//'",
			path:     "//",
			expected: []string{},
		},
		{
			name:     "Path '/tmp'",
			path:     "/tmp",
			expected: []string{"tmp"},
		},
		{
			name:     "Path './tmp'",
			path:     "./tmp",
			expected: []string{"tmp"},
		},
		{
			name:     "Path '../tmp'",
			path:     "../tmp",
			expected: []string{"tmp"},
		},
		{
			name:     "Path './tmp/log'",
			path:     "./tmp/log",
			expected: []string{"tmp", "log"},
		},
		{
			name:     "Path 'tmp/log/'",
			path:     "tmp/log/",
			expected: []string{"tmp", "log"},
		},
		{
			name:     "Path 'tmp/log/soft.log'",
			path:     "tmp/log/soft.log",
			expected: []string{"tmp", "log", "soft.log"},
		},
		{
			name:     "Path '../tmp/log/soft.log/'",
			path:     "../tmp/log/soft.log/",
			expected: []string{"tmp", "log", "soft.log"},
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				act := common.SplitPath(tc.path)
				assertSlices(t, act, tc.expected)
			})
	}
}

func assertSlices(t *testing.T, got, want []string) {
	if reflect.DeepEqual(got, want) {
		return
	}

	t.Errorf("got %v want %v", got, want)
}
