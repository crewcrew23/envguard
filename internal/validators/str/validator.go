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

func (v *validator) Length(min, max int) interfaces.StringValidator {
	if v.err == nil && (len(v.value) < min || len(v.value) > max) {
		v.err = errs.ErrStrLength
	}
	return v
}

func (v *validator) Min(min int) interfaces.StringValidator {
	if v.err == nil && len(v.value) < min {
		v.err = errs.ErrStrMin
	}
	return v
}

func (v *validator) Max(max int) interfaces.StringValidator {
	if v.err == nil && len(v.value) > max {
		v.err = errs.ErrStrMax
	}
	return v
}

func (v *validator) IsAlpha() interfaces.StringValidator {
	if v.err != nil {
		return v
	}

	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if !isAlpha(v.value) {
		v.err = fmt.Errorf("%s is not Alpha", v.value)
	}

	return v
}

func (v *validator) IsAlphanumeric() interfaces.StringValidator {
	if v.err != nil {
		return v
	}

	isAlphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString
	if !isAlphanumeric(v.value) {
		v.err = fmt.Errorf("%s is not Alphanumeric", v.value)
	}

	return v
}

func (v *validator) NotEmpty() interfaces.StringValidator {
	if v.err == nil && len(v.value) == 0 {
		v.err = errs.ErrStrEmpty
	}
	return v
}

func (v *validator) NotBlank() interfaces.StringValidator {
	if v.err == nil && strings.TrimSpace(v.value) == "" {
		v.err = errs.ErrStrBlank
	}
	return v
}

func (v *validator) HasPrefix(prefix string) interfaces.StringValidator {
	if v.err == nil && !strings.HasPrefix(v.value, prefix) {
		v.err = fmt.Errorf("%s doesnt has prefix %s", v.value, prefix)
	}

	return v
}

func (v *validator) HasSuffix(suffix string) interfaces.StringValidator {
	if v.err == nil && !strings.HasSuffix(v.value, suffix) {
		v.err = fmt.Errorf("%s doesnt has suffix %s", v.value, suffix)
	}

	return v
}

func (v *validator) Email() interfaces.StringValidator {
	if v.err == nil && isEmailValid(v.value) {
		v.err = errs.ErrNotEmail
	}
	return v
}

func (v *validator) UUID() interfaces.StringValidator {
	if v.err != nil {
		return v
	}

	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	if !r.MatchString(v.value) {
		v.err = fmt.Errorf("%s is not UUID", v.value)
	}

	return v
}

func (v *validator) MatchRegex(pattern string) interfaces.StringValidator {
	if v.err == nil {
		regexpPattern := regexp.MustCompile(pattern)
		if !regexpPattern.MatchString(v.value) {
			v.err = errs.ErrRegexp
		}
	}
	return v
}

func (v *validator) Contains(substr string) interfaces.StringValidator {
	if v.err == nil && !strings.Contains(v.value, substr) {
		v.err = errs.ErrNotContains
	}
	return v
}

func (v *validator) NotContains(forbidden string) interfaces.StringValidator {
	if v.err == nil && strings.Contains(v.value, forbidden) {
		v.err = errs.ErrContains
	}
	return v
}

func (v *validator) Custom(fn func(string) bool, errmsg string) interfaces.StringValidator {
	if v.err == nil && !fn(v.value) {
		v.err = errors.New(errmsg)
	}
	return v
}

func (v *validator) Validate() error {
	return v.err
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
