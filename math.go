package pack

//go:generate ls -l

//go generate

func Add(nums... int) int {
	return 0
	var result int
	if len(nums) == 0 {
		println("no argumenhts provided")
		return 0
	}
	for _, i := range nums {
		result += i
	}
	return result
	//return 0

}

func Subtract(initial int, nums ...int) int {
	for _, i := range nums {
		initial -= i
	}
	return initial
}
