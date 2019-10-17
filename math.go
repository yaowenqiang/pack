package pack

//go:generate ls -l

//go generate

func Add(nums... int) int {
	var result int
	for _, i := range nums {
		result += i
	}
	return result
	//return 0

}
