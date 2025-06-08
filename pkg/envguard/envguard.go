package envguard

import (
	"github.com/crewcrew23/envguard/internal/envparser"
	"github.com/crewcrew23/envguard/internal/envtypes"
	"github.com/crewcrew23/envguard/internal/validators/integer"
	"github.com/crewcrew23/envguard/internal/validators/interfaces"
)

type EnvMap struct {
	envMap *envtypes.EnvMap
}

func (em *EnvMap) Get(key string) interfaces.Validator {
	return &validatorAdapter{value: em.envMap.Get(key)}
}

func ParseFile(filePath string) (*EnvMap, error) {
	envMap, err := envparser.ParseFile(filePath)
	if err != nil {
		return nil, err
	}
	return &EnvMap{envMap: envMap}, nil
}

type validatorAdapter struct {
	value *envtypes.EnvValue
}

func (va *validatorAdapter) Integer() interfaces.IntegerValidator {
	return integer.New(va.value)
}
