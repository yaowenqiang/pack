package pack

import (
	"testing"
	"time"
)

//go test -timeout 2s -v

func TestCanAddMembers(t *testing.T) {
	if testing.Short() {
		//go test -short
		t.Skip("Skipping long tests")	
	}
	t.Parallel()
	result := Add(1,2)
	time.Sleep(3 + time.Second)
	if result != 3 {
		//t.Log("Failed to add one and two")
		t.Fatal("Failed to add one and two")
		//t.Fail()
		//t.FailNow()
	}

	result = Add(1,2, 3,4)
	if result != 10 {
		t.Error("Failed to add more than two numbers")
	}

}


func TestSubtractNumbers(t *testing.T) {
	t.Parallel()
	result := Subtract(10,5)
	time.Sleep(1 + time.Second)
	if result != 5 {
		t.Log("Failed to substrct two numbers properly")
		t.Fail()
	}
}

func TestCanMltiplyNmers(t *testing.T) {
	//t.Skip("Not implemented yet!")
	if testing.Verbose() {
		t.Skip("Not implemented yet")
	}
}

