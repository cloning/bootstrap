package service_test

import (
	"../../../src/app/service"
	"testing"
)

func TestName(t *testing.T) {
	s := service.NewService("Test2")
	u := s.GetUser()
	if u.Name != "Test2" {
		t.Errorf("Expected %s, but was %s", "Test2", u.Name)
	}
}
