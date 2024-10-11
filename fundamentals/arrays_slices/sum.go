package sum

func Sum(numbers [5]int) int{
	sum := 0
	//for i := 0; i<5;i++{
	//	sum += numbers[i]
	//}
	// same as python: index and value
	for _, number := range numbers {
		sum += number
	}
	return sum
}