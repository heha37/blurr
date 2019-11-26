package suitemaker

import (
	"reflect"
)

var (
	EnvTypeMap = make(map[string]Environment)
)

type Environment interface {
	Setup()
	TearDown()
}

func RegisterEnv(vals ...interface{}) {
	for _, val := range vals {
		env, _ := val.(Environment)
		name := reflect.TypeOf(val).Elem().Name()
		EnvTypeMap[name] = env
	}
}
