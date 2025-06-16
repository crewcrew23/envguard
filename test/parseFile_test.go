package test

import (
	"testing"

	"github.com/crewcrew23/envguard/pkg/envguard"
)

func TestLoadFileValid(t *testing.T) {
	_, err := envguard.ParseFile("./.env")
	if err != nil {
		t.Fail()
	}
}

func TestLoadFileInvalid(t *testing.T) {
	_, err := envguard.ParseFile("./abc")
	if err == nil {
		t.Fail()
	}
}
