package arraysnslices

func Sum(n []int) int {
	sum := 0

	for _, v := range n {
		sum += v

	}

	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	var arr []int

	for _, number := range numbersToSum {
		arr = append(arr, Sum(number))
	}

	return arr

}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	var arr []int

	for _, numbers := range numbersToSum {

		if len(numbers) == 0 {
			arr = append(arr, 0)

		} else {

			tail := numbers[1:]
			arr = append(arr, Sum(tail))
		}

	}

	return arr
}
