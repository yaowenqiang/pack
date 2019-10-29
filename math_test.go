package pack

import (
	"testing"
	_ "testing/quick"
	"time"
	 "os"
)

var addTable = []struct {
	in []int
	out int
} {
	{[]int{1,2},3},
	{[]int{1,2,3, 4}, 10},
}


// custom runner


func TestMain(m *testing.M) {
	println("Tests are about to run!")
	result := m.Run()
	println("Tests done executing")
	os.Exit(result)
}


//go test -timeout 2s -v

func TestCanAddMembers(t *testing.T) {
	if testing.Short() {
		//go test -short
		t.Skip("Skipping long tests")	
	}


	for _, entry := range addTable {
		result := Add(entry.in...)
		if result != entry.out {
			t.Error("Failed to add numbers ad expected!")
		}
	}
	/*
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
	*/

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

/*
func TestOddMultipleOfThree(t *testing.T) {
	f := func(x int) bool {
		y := OldMultipleOfThree(x)
		return y%2 == 1 && y%3 == 0
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
*/

