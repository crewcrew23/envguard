package interfaces

type Validator interface {
	Integer() IntegerValidator
	String() StringValidator
	URL() URLValidator
	IP() IPValidator
}

type IntegerValidator interface {
	Min(int) IntegerValidator
	Max(int) IntegerValidator
	Between(int, int) IntegerValidator

	Even() IntegerValidator
	Odd() IntegerValidator

	Contains(...int) IntegerValidator
	NotContains(...int) IntegerValidator

	Positive() IntegerValidator
	Negative() IntegerValidator
	NonZero() IntegerValidator

	DivisibleBy(divisor int) IntegerValidator

	Custom(func(int) bool, string) IntegerValidator
	Validate() error
}

type StringValidator interface {
	Length(min, max int) StringValidator
	Min(min int) StringValidator
	Max(max int) StringValidator

	IsAlpha() StringValidator
	IsAlphanumeric() StringValidator

	NotEmpty() StringValidator
	NotBlank() StringValidator

	HasPrefix(string) StringValidator
	HasSuffix(string) StringValidator

	Email() StringValidator
	UUID() StringValidator
	MatchRegexp(pattern string) StringValidator

	Contains(substr string) StringValidator
	NotContains(forbidden string) StringValidator

	Custom(func(string) bool, string) StringValidator
	Validate() error
}

type URLValidator interface {
	Scheme(schemes ...string) URLValidator
	Host(hosts ...string) URLValidator
	Port(ports ...string) URLValidator
	Custom(func(string) bool, string) URLValidator
	Validate() error
}

type IPValidator interface {
	V4() IPValidator
	V6() IPValidator
	Custom(func(string) bool, string) IPValidator
	Validate() error
}
