package interfaces

type Validator interface {
	Integer() IntegerValidator
	String() StringValidator
}

type IntegerValidator interface {
	Min(min int) IntegerValidator
	Max(max int) IntegerValidator

	Even() IntegerValidator
	Odd() IntegerValidator

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

	NotEmpty() StringValidator
	NotBlank() StringValidator

	Email() StringValidator
	MatchRegex(pattern string) StringValidator
	URL() StringValidator

	Contains(substr string) StringValidator
	NotContains(forbidden string) StringValidator

	Custom(func(string) bool, string) StringValidator
	Validate() error
}
