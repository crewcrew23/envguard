package test

import (
	"testing"

	"github.com/crewcrew23/envguard/pkg/envguard"
)

func TestStringLengthValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("STR").String().Length(1, 10).Validate()
	if err != nil {
		t.Fail()
	}
}

func TestStringLengthInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("STR").String().Length(20, 30).Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringMinhValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("STR").String().Min(1).Validate()
	if err != nil {
		t.Fail()
	}
}

func TestStringMinInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("STR").String().Min(10).Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringMaxhValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("STR").String().Max(10).Validate()
	if err != nil {
		t.Fail()
	}
}

func TestStringMaxInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("STR").String().Max(1).Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringIsAlphahValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("ALPHA").String().IsAlpha().Validate()
	if err != nil {
		t.Fail()
	}
}

func TestStringIsAlphaInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NOTALPHA").String().IsAlpha().Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringIsAlphahNumericValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("ALPHANUM").String().IsAlphanumeric().Validate()
	if err != nil {
		t.Fail()
	}
}

func TestStringIsAlphaNumericInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NOTALPHA").String().IsAlphanumeric().Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringNotEmptyValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NOTALPHA").String().NotBlank().Validate()
	if err != nil {
		t.Fail()
	}
}

func TestStringIsNotEmptyInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("EMPPTY").String().NotEmpty().Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringNotBlankValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NOTALPHA").String().NotBlank().Validate()
	if err != nil {
		t.Fail()
	}
}

func TestStringIsNotBlankInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("BLANK").String().NotBlank().Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringHasPrefixValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NOTCONTAINS").String().HasPrefix("go").Validate()
	if err != nil {
		t.Fail()
	}
}

func TestStringIsHasPrefixInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NOTCONTAINS").String().HasPrefix("he").Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringHasSuffixValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NOTCONTAINS").String().HasSuffix("ng").Validate()
	if err != nil {
		t.Fail()
	}
}

func TestStringIsHasSuffixInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NOTCONTAINS").String().HasSuffix("go").Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringEmailValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("EMAIL").String().Email().Validate()
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestStringIsEmailInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NOTEMAIL").String().Email().Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringUUIDValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("UUID").String().UUID().Validate()
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestStringUUIDInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NOTALPHA").String().UUID().Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringMatchRegexpValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("ALPHA").String().MatchRegexp("^[A-Za-z]+$").Validate()
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestStringMatchRegexpInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("NOTALPHA").String().MatchRegexp("^[A-Za-z]+$").Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringContainsValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("CONTAINS").String().Contains("he").Validate()
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestStringContainsInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("CONTAINS").String().Contains("go").Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringNotContainsValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("CONTAINS").String().NotContains("go").Validate()
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestStringNotContainsInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("CONTAINS").String().NotContains("he").Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringCustom(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("CONTAINS").String().Custom(func(s string) bool {
		return len(s) > 1
	}, "err msg").Validate()

	if err != nil {
		t.Fail()
	}
}
