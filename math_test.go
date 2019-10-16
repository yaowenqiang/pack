package pack

import (
	"testing"
)

func TestCanAddMembers(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Log("Failed to add one and two")
		t.Fail()
	}
}
