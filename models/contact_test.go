package models_test

import (
	"contacts/models"
	"testing"
)

func TestValidatePass(t *testing.T) {
	contact := &models.Contact{Email: "Jitenp@outlook.com", Mobile: "9618558500", Name: "Jiten"}
	err := contact.Validate()
	if err != nil {
		t.Error("Validate must not  return an error")
		t.Fail()
	}

}
func TestValidateFail(t *testing.T) {
	contact := &models.Contact{}
	err := contact.Validate()
	if err == nil {
		t.Error("Validate must return an error")
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	actual := models.Add(12, 13)
	expected := 25
	if actual != expected {
		t.Errorf("The result must be %v, but returned %v", expected, actual)
	}
}
