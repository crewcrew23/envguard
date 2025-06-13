package ip

import (
	"errors"
	"fmt"
	"net"

	"github.com/crewcrew23/envguard/internal/envtypes"
	"github.com/crewcrew23/envguard/internal/validators/interfaces"
)

type validator struct {
	value string
	err   error
}

func New(provide envtypes.ValueProvider) interfaces.IPValidator {
	if provide.Error() != nil {
		return &validator{err: provide.Error()}
	}
	return &validator{value: provide.Value()}
}

func (v *validator) V4() interfaces.IPValidator {
	if v.err != nil {
		return v
	}

	ip := net.ParseIP(v.value)
	if ip == nil {
		v.err = fmt.Errorf("'%s' is not a valid IP address", v.value)
		return v
	}

	if ip.To4() == nil {
		v.err = fmt.Errorf("'%s' is not an IPv4 address", v.value)
	}

	return v
}

func (v *validator) V6() interfaces.IPValidator {
	if v.err != nil {
		return v
	}

	ip := net.ParseIP(v.value)
	if ip == nil {
		v.err = fmt.Errorf("'%s' is not a valid IP address", v.value)
		return v
	}

	if ip.To4() != nil || ip.To16() == nil {
		v.err = fmt.Errorf("'%s' is not an IPv6 address", v.value)
	}

	return v
}

func (v *validator) Custom(fn func(string) bool, errmsg string) interfaces.IPValidator {
	if v.err == nil && !fn(v.value) {
		v.err = errors.New(errmsg)
	}

	return v
}

func (v *validator) Validate() error {
	return v.err
}
