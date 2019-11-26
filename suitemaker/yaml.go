package suitemaker

import (
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

func ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

func LoadYAML(name string) (suite TestSuite, err error) {
	file, err := ioutil.ReadFile(name)
	s := new(TestSuite)
	err = yaml.Unmarshal([]byte(file), &s)
	suite = *s
	return
}
