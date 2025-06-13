package url

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/crewcrew23/envguard/internal/envtypes"
	"github.com/crewcrew23/envguard/internal/validators/interfaces"
)

type validator struct {
	value string
	err   error
}

func New(provide envtypes.ValueProvider) interfaces.URLValidator {
	if provide.Error() != nil {
		return &validator{err: provide.Error()}
	}
	return &validator{value: provide.Value()}
}

func (v *validator) Scheme(schemes ...string) interfaces.URLValidator {
	if v.err != nil {
		return v
	}

	if len(schemes) == 0 {
		v.err = fmt.Errorf("no schemes provided for validation")
		return v
	}

	parsedURL, err := url.Parse(v.value)
	if err != nil {
		v.err = fmt.Errorf("invalid URL: %w", err)
		return v
	}

	urlScheme := strings.ToLower(parsedURL.Scheme)

	found := false
	for _, s := range schemes {
		if strings.ToLower(s) == urlScheme {
			found = true
			break
		}
	}

	if !found {
		v.err = fmt.Errorf("invalid scheme '%s', expected one of: %v",
			parsedURL.Scheme, schemes)
	}

	return v
}

func (v *validator) Host(hosts ...string) interfaces.URLValidator {
	if v.err != nil {
		return v
	}

	if len(hosts) == 0 {
		v.err = fmt.Errorf("no hosts provided for validation")
	}

	parsedURL, err := url.Parse(v.value)
	if err != nil {
		v.err = fmt.Errorf("invalid URL: %w", err)
		return v
	}

	host := strings.ToLower(parsedURL.Host)

	found := false
	for _, s := range hosts {
		if strings.ToLower(s) == host {
			found = true
			break
		}
	}

	if !found {
		v.err = fmt.Errorf("invalid host '%s', expected one of: %v",
			parsedURL.Host, hosts)
	}

	return v
}

func (v *validator) Port(ports ...string) interfaces.URLValidator {
	if v.err != nil {
		return v
	}

	if len(ports) == 0 {
		v.err = fmt.Errorf("no ports provided for validation")
	}

	parsedURL, err := url.Parse(v.value)
	if err != nil {
		v.err = fmt.Errorf("invalid URL: %w", err)
		return v
	}

	var port string
	switch {
	case parsedURL.Port() != "":
		port = parsedURL.Port()
	case parsedURL.Opaque != "":
		port = parsedURL.Opaque
	}

	found := false
	for _, p := range ports {
		if p == port {
			found = true
			break
		}
	}

	if !found {
		v.err = fmt.Errorf("invalid port '%s', expected one of: %v",
			parsedURL.Port(), ports)
	}

	return v
}

func (v *validator) Custom(fn func(string) bool, errmsg string) interfaces.URLValidator {
	if v.err == nil && !fn(v.value) {
		v.err = errors.New(errmsg)
	}

	return v
}

func (v *validator) Validate() error {
	return v.err
}
