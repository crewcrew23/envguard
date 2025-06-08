package envtypes

type ValueProvider interface {
	Value() string
	Error() error
}

type EnvValue struct {
	value string
	err   error
}

func (ev *EnvValue) Value() string {
	return ev.value
}

func (ev *EnvValue) Error() error {
	return ev.err
}
