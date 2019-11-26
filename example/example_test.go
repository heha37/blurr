package example

import (
	"testing"

	"github.com/shellpickup/blurr/suitemaker"
)

func TestExample(t *testing.T) {
	suitemaker.RegisterEnv(new(RedisEnv))
	suitemaker.BuildTests(t, "blurrs")
}
