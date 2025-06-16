package test

import (
	"testing"

	"github.com/crewcrew23/envguard/pkg/envguard"
)

func TestIPV4Valid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("IP").IP().V4().Validate()
	if err != nil {
		t.Fail()
	}
}

func TestIPV4Invalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("PORT").IP().V4().Validate()
	if err == nil {
		t.Fail()
	}
}

func TestIPV6Valid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("IPV6").IP().V6().Validate()
	if err != nil {
		t.Fail()
	}
}

func TestIPV6Invalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("IP").IP().V6().Validate()
	if err == nil {
		t.Fail()
	}
}

func TestIPCustom(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("IP").IP().Custom(func(s string) bool {
		return true
	}, "err msg").Validate()

	if err != nil {
		t.Fail()
	}
}
