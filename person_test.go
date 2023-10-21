package person

import (
	"testing"
)

func TestNew(t *testing.T){
	expected := Person{"Tommy", 42}
	actual := New("Tommy", 42)
	if actual != expected {
		t.Error("New should correctly instantiate a Person struct if provided with valid data. Got: ", actual)
	}
}