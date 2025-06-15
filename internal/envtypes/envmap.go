package envtypes

import "github.com/crewcrew23/envguard/internal/errs"

type EnvMap struct {
	values map[string]*EnvValue
}

func NewEnvMap() *EnvMap {
	return &EnvMap{
		values: make(map[string]*EnvValue),
	}
}

func (e *EnvMap) Get(key string) *EnvValue {
	if val, exists := e.values[key]; exists {
		return val
	}
	return &EnvValue{err: errs.ErrKeyNotFound}
}

func (e *EnvMap) Set(key, value string) {
	e.values[key] = &EnvValue{value: value}
}
