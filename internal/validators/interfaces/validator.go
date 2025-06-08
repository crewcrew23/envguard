package interfaces

type Validator interface {
	Integer() IntegerValidator
}

type IntegerValidator interface {
	Min(min int) IntegerValidator
	Max(max int) IntegerValidator
	Validate() error
}
