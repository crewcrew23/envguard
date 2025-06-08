package integer

import (
	"strconv"

	"github.com/crewcrew23/envguard/internal/envtypes"
	"github.com/crewcrew23/envguard/internal/errs"
	"github.com/crewcrew23/envguard/internal/validators/interfaces"
)

type validator struct {
	value int
	err   error
}

func New(provider envtypes.ValueProvider) interfaces.IntegerValidator {
	if provider.Error() != nil {
		return &validator{err: provider.Error()}
	}

	num, err := strconv.Atoi(provider.Value())
	if err != nil {
		return &validator{err: err}
	}

	return &validator{
		value: num,
	}
}

func (v *validator) Min(min int) interfaces.IntegerValidator {
	if v.err == nil && v.value < min {
		v.err = errs.ErrMinValue
	}
	return v
}

func (v *validator) Max(max int) interfaces.IntegerValidator {
	if v.err == nil && v.value > max {
		v.err = errs.ErrMaxValue
	}
	return v
}

func (v *validator) Validate() error {
	return v.err
}
