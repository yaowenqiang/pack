package pack
import (
	"testing"
)


var table = []struct {
	in string
	out *NumberData
	err string
} {
   {
	    in: "not a number",
		out: &NumberData {
			value: 0,
			 isNumeric: false,
			 isInteger: false,
			 isNegative: false,
		},
		err: "Did not evaluate non-number correctly",
	}, 
   {
	    in: "not a number",
		out: &NumberData {
			value: 0,
			 isNumeric: false,
			 isInteger: false,
			 isNegative: false,
		},
		err: "Did not evaluate non-number correctly",
	}, 
   {
	    in: "not a number",
		out: &NumberData {
			value: 0,
			 isNumeric: false,
			 isInteger: false,
			 isNegative: false,
		},
		err: "Did not evaluate non-number correctly",
	}, 
	{
		in: "42.2",
		out: &NumberData {
			 value: 42.2,
			 isNumeric: true,
			 isInteger: false,
			 isNegative: false,
		},
		err: "Did not evaluate positive, non-integer number correctly",
	},
}


func TestEvaluateNotNumberCorrectly(t *testing.T) {
	for _, entry := range table {
		candidate := entry.in

		result := NumberEvaluator(candidate)

		if result.isNumeric != entry.out.isNumeric ||
			result.value != entry.out.value ||
			result.isNegative != entry.out.isNegative ||
			result.isInteger != entry.out.isInteger {
			t.Log(result)
			t.Log(entry.out)
			t.Error(entry.err)
		}
	}
}
