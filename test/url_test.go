package test

import (
	"net/url"
	"testing"

	"github.com/crewcrew23/envguard/pkg/envguard"
)

func TestUrlSchemeValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("CORRECTSCHEME").URL().Scheme("https").Validate()
	if err != nil {
		t.Fail()
	}
}

func TestUrlSchemeInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("CORRECTSCHEME").URL().Scheme("ws").Validate()
	if err == nil {
		t.Fail()
	}
}

func TestUrlHostValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("CORRECTSCHEME").URL().Host("google.com").Validate()
	if err != nil {
		t.Fail()
	}
}

func TestUrlHostInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("CORRECTSCHEME").URL().Host("amazon.com").Validate()
	if err == nil {
		t.Fail()
	}
}

func TestUrlPortValid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("PORT").URL().Port("8080").Validate()
	if err != nil {
		t.Fail()
	}
}

func TestUrlPortInvalid(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("PORT").URL().Port("5432").Validate()
	if err == nil {
		t.Fail()
	}
}

func TestUrlCustom(t *testing.T) {
	file, _ := envguard.ParseFile("./.env")

	err := file.Get("PORT").URL().Custom(func(s string) bool {
		_, err := url.Parse(s)
		return err == nil

	}, "err msg").Validate()
	if err != nil {
		t.Fail()
	}
}
