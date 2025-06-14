package str

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/crewcrew23/envguard/internal/envtypes"
	"github.com/crewcrew23/envguard/internal/validators/interfaces"
	"github.com/crewcrew23/envguard/pkg/errs"
)

type validator struct {
	value string
	err   error
}

func New(provide envtypes.ValueProvider) interfaces.StringValidator {
	if provide.Error() != nil {
		return &validator{err: provide.Error()}
	}
	return &validator{value: provide.Value()}
}

func (sv *validator) Length(min, max int) interfaces.StringValidator {
	if sv.err == nil && (len(sv.value) < min || len(sv.value) > max) {
		sv.err = errs.ErrStrLength
	}
	return sv
}

func (sv *validator) Min(min int) interfaces.StringValidator {
	if sv.err == nil && len(sv.value) < min {
		sv.err = errs.ErrStrMin
	}
	return sv
}

func (sv *validator) Max(max int) interfaces.StringValidator {
	if sv.err == nil && len(sv.value) > max {
		sv.err = errs.ErrStrMax
	}
	return sv
}

func (sv *validator) IsAlpha() interfaces.StringValidator {
	if sv.err != nil {
		return sv
	}

	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if !isAlpha(sv.value) {
		sv.err = fmt.Errorf("%s is not Alpha", sv.value)
	}

	return sv
}

func (sv *validator) IsAlphanumeric() interfaces.StringValidator {
	if sv.err != nil {
		return sv
	}

	isAlphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString
	if !isAlphanumeric(sv.value) {
		sv.err = fmt.Errorf("%s is not Alphanumeric", sv.value)
	}

	return sv
}

func (sv *validator) NotEmpty() interfaces.StringValidator {
	if sv.err == nil && len(sv.value) == 0 {
		sv.err = errs.ErrStrEmpty
	}
	return sv
}

func (sv *validator) NotBlank() interfaces.StringValidator {
	if sv.err == nil && strings.TrimSpace(sv.value) == "" {
		sv.err = errs.ErrStrBlank
	}
	return sv
}

func (sv *validator) Email() interfaces.StringValidator {
	if sv.err == nil && isEmailValid(sv.value) {
		sv.err = errs.ErrNotEmail
	}
	return sv
}

func (sv *validator) MatchRegex(pattern string) interfaces.StringValidator {
	if sv.err == nil {
		regexpPattern := regexp.MustCompile(pattern)
		if !regexpPattern.MatchString(sv.value) {
			sv.err = errs.ErrRegexp
		}
	}
	return sv
}

func (sv *validator) Contains(substr string) interfaces.StringValidator {
	if sv.err == nil && !strings.Contains(sv.value, substr) {
		sv.err = errs.ErrNotContains
	}
	return sv
}

func (sv *validator) NotContains(forbidden string) interfaces.StringValidator {
	if sv.err == nil && strings.Contains(sv.value, forbidden) {
		sv.err = errs.ErrContains
	}
	return sv
}

func (sv *validator) Custom(fn func(string) bool, errmsg string) interfaces.StringValidator {
	if sv.err == nil && !fn(sv.value) {
		sv.err = errors.New(errmsg)
	}
	return sv
}

func (sv *validator) Validate() error {
	return sv.err
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
