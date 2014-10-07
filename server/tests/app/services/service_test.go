package service_test

import (
	"../../../src/app/services"
	"testing"
)

func TestName(t *testing.T) {
	s := services.NewService("Test2")
	u := s.GetUser()
	if u.Name != "Test2" {
		t.Errorf("Expected %s, but was %s", "Test2", u.Name)
	}
}
