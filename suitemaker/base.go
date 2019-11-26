package suitemaker

import (
	"encoding/json"
	"reflect"

	"github.com/shellpickup/blurr/target"
	"github.com/shellpickup/blurr/utils"
	"github.com/stretchr/testify/suite"
)

type TestCase struct {
	Name     string                 `yaml:"name"`
	Method   string                 `yaml:"method"`
	Path     string                 `yaml:"path"`
	Headers  map[string]string      `yaml:"headers"`
	Cookies  map[string]string      `yaml:"cookies"`
	Request  string                 `yaml:"request"`
	Response map[string]interface{} `yaml:"response"`
}

type TestSuite struct {
	suite.Suite

	Intercepter target.Intercept

	Envs      []string   `yaml:"envs"`
	TestCases []TestCase `yaml:"test_cases"`
}

func (t *TestSuite) SetupTest() {
	for _, env := range t.Envs {
		EnvTypeMap[env].Setup()
	}
}

func (t *TestSuite) TearDownTest() {
	for _, env := range t.Envs {
		EnvTypeMap[env].TearDown()
	}
}

func (t *TestSuite) TestAllTestCases() {
	tests := make([]func(), 0)
	for _, tc := range t.TestCases {
		f := func(tc TestCase) func() {
			return func() {
				var data []byte
				if t.Intercepter == nil {
					data, _ = utils.CallAPI(tc.Method, tc.Path, tc.Request, tc.Headers, tc.Cookies)
				} else {
					rw := target.CallAPI(t.Intercepter, tc.Method, tc.Path, tc.Request, tc.Headers)
					data = rw.Body.Bytes()
				}
				var jsonData, comparedData interface{}

				if len(data) > 0 {
					json.Unmarshal(data, &jsonData)
				}

				json.Unmarshal([]byte(tc.Response["data"].(string)), &comparedData)

				ok := reflect.DeepEqual(jsonData, comparedData)
				if ok {
					t.T().Log("SUCCESS:", tc.Name)
				} else {
					t.T().Log("FAILED:", tc.Name)
				}
			}
		}(tc)
		tests = append(tests, f)
	}

	for _, test := range tests {
		test()
	}
}
