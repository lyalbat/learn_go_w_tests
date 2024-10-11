package sum

func Sum(numbers []int) int{
	sum := 0
	//for i := 0; i<5;i++{
	//	sum += numbers[i]
	//}
	// similar to python dics range: index and value, but no need to specify len(numbers)
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//variadic function - reference: gobyexample.com/variadic-functions
func SumAll(numbersToSum ...[]int) []int{
	//sizeNumbers := len(numbersToSum)
	//make can create slices with specific capacities
	//in this case it takes sizeNumbers as the cap
	//sums := make([]int,sizeNumbers)

	//for i, numbers := range numbersToSum {
	//	sums[i] = Sum(numbers)
	//}
	//if you tried to run append with the original make you would get:
	//[0 0 3 9] instead of [3 9], because append does not overwrite
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums,Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int{
	var sums []int
	
	for _,numbers := range numbersToSum {

		if len(numbers) == 0 {
			sums = append(sums,Sum(numbers))
		} else{
			tail := numbers[1:]
			sums = append(sums,Sum(tail))
		}
	}
	return sums
}