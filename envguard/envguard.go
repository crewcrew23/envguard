package envguard

import (
	"github.com/crewcrew23/envguard/internal/envparser"
	"github.com/crewcrew23/envguard/internal/envtypes"
	"github.com/crewcrew23/envguard/internal/validators/integer"
	"github.com/crewcrew23/envguard/internal/validators/interfaces"
	"github.com/crewcrew23/envguard/internal/validators/ip"
	"github.com/crewcrew23/envguard/internal/validators/str"
	"github.com/crewcrew23/envguard/internal/validators/url"
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

func (va *validatorAdapter) String() interfaces.StringValidator {
	return str.New(va.value)
}

func (va *validatorAdapter) URL() interfaces.URLValidator {
	return url.New(va.value)
}

func (va *validatorAdapter) IP() interfaces.IPValidator {
	return ip.New(va.value)

}
