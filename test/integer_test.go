package test

import (
	"testing"

	"github.com/crewcrew23/envguard/pkg/envguard"
)

func TestIntegerMinValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("INT").Integer().Min(1).Validate()
	if err != nil {
		t.Fail()
	}
}

func TestIntegerMinInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("INT").Integer().Min(5).Validate()
	if err == nil {
		t.Fail()
	}
}

func TestIntegerMaxValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("INT").Integer().Max(10).Validate()
	if err != nil {
		t.Fail()
	}
}

func TestIntegerMaxInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("INT").Integer().Max(0).Validate()
	if err == nil {
		t.Fail()
	}
}

func TestIntegerBetweenValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("INT").Integer().Between(0, 5).Validate()
	if err != nil {
		t.Fail()
	}
}

func TestIntegerBetweenInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("INT").Integer().Between(6, 9).Validate()
	if err == nil {
		t.Fail()
	}
}

func TestIntegerEvenValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("EVEN").Integer().Even().Validate()
	if err != nil {
		t.Fail()
	}
}

func TestIntegerEvenInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("ODD").Integer().Even().Validate()
	if err == nil {
		t.Fail()
	}
}

func TestIntegerOddValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("ODD").Integer().Odd().Validate()
	if err != nil {
		t.Fail()
	}
}

func TestIntegerOddInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("EVEN").Integer().Odd().Validate()
	if err == nil {
		t.Fail()
	}
}

func TestIntegerContainsValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("INT").Integer().Contains(1, 2, 3).Validate()
	if err != nil {
		t.Fail()
	}
}

func TestIntegerContainsInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("INT").Integer().Contains(1, 3).Validate()
	if err == nil {
		t.Fail()
	}
}

func TestIntegerNotContainsValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("INT").Integer().NotContains(1, 3).Validate()
	if err != nil {
		t.Fail()
	}
}

func TestIntegerNotContainsInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("INT").Integer().NotContains(1, 2, 3).Validate()
	if err == nil {
		t.Fail()
	}
}

func TestIntegerPositiveValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("INT").Integer().Positive().Validate()
	if err != nil {
		t.Fail()
	}
}

func TestIntegerPositiveInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NEGATIVE").Integer().Positive().Validate()
	if err == nil {
		t.Fail()
	}
}

func TestIntegerNegativeValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NEGATIVE").Integer().Negative().Validate()
	if err != nil {
		t.Fail()
	}
}

func TestIntegerNegativeInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("INT").Integer().Negative().Validate()
	if err == nil {
		t.Fail()
	}
}

func TestIntegerNonZeroValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NEGATIVE").Integer().NonZero().Validate()
	if err != nil {
		t.Fail()
	}
}

func TestIntegerNonZeroInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("ZERO").Integer().NonZero().Validate()
	if err == nil {
		t.Fail()
	}
}

func TestIntegerDivisibleByValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("DEVIDE").Integer().DivisibleBy(2).Validate()
	if err != nil {
		t.Fail()
	}
}

func TestIntegerDivisibleByInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("DEVIDE").Integer().DivisibleBy(3).Validate()
	if err == nil {
		t.Fail()
	}
}

func TestIntegerCustom(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("DEVIDE").Integer().Custom(func(i int) bool {
		return i%2 == 0

	}, "err").Validate()

	if err != nil {
		t.Fail()
	}
}
