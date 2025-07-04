package integer

import (
	"errors"
	"fmt"
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

func (v *validator) Between(min, max int) interfaces.IntegerValidator {
	if v.err != nil {
		return v
	}

	if v.value < min {
		v.err = fmt.Errorf("%d is less than min", v.value)
	}

	if v.value > max {
		v.err = fmt.Errorf("%d is greater than max", v.value)
	}

	return v
}

func (v *validator) Even() interfaces.IntegerValidator {
	if v.err == nil && v.value%2 != 0 {
		v.err = errs.ErrNotEven
	}
	return v
}
func (v *validator) Odd() interfaces.IntegerValidator {
	if v.err == nil && v.value%2 == 0 {
		v.err = errs.ErrNotOdd
	}
	return v
}

func (v *validator) Contains(alowed ...int) interfaces.IntegerValidator {
	if v.err != nil {
		return v
	}

	find := false
	for _, i := range alowed {
		if i == v.value {
			find = true
			break
		}
	}

	if !find {
		v.err = fmt.Errorf("%d is not contains", v.value)
	}

	return v
}
func (v *validator) NotContains(disallowed ...int) interfaces.IntegerValidator {
	if v.err != nil {
		return v
	}

	find := false
	for _, i := range disallowed {
		if i == v.value {
			find = true
			break
		}
	}

	if find {
		v.err = fmt.Errorf("%d is disallowed", v.value)
	}

	return v
}

func (v *validator) Positive() interfaces.IntegerValidator {
	if v.err == nil && v.value < 0 {
		v.err = errs.ErrNotPositive
	}
	return v
}

func (v *validator) Negative() interfaces.IntegerValidator {
	if v.err == nil && v.value > -1 {
		v.err = errs.ErrNotNegative
	}
	return v
}

func (v *validator) NonZero() interfaces.IntegerValidator {
	if v.err == nil && v.value == 0 {
		v.err = errs.ErrZero
	}
	return v
}

func (v *validator) DivisibleBy(divisor int) interfaces.IntegerValidator {
	if v.err == nil && v.value%divisor != 0 {
		v.err = fmt.Errorf("value %d not divide by %d", v.value, divisor)
	}
	return v
}

func (v *validator) Custom(fn func(int) bool, errmasg string) interfaces.IntegerValidator {
	if v.err == nil && !fn(v.value) {
		v.err = errors.New(errmasg)
	}
	return v
}

func (v *validator) Validate() error {
	return v.err
}
