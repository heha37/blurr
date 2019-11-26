package suitemaker

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

func BuildTests(t *testing.T, path string) {
	suiteList := make([]*TestSuite, 0)
	fileList, _ := ReadDir(path)
	for _, file := range fileList {
		s := new(TestSuite)
		ts, _ := LoadYAML(filepath.Join(path, file.Name()))
		s.Envs = ts.Envs
		s.TestCases = ts.TestCases

		suiteList = append(suiteList, s)
	}

	for _, s := range suiteList {
		suite.Run(t, s)
	}
}
