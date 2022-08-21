package iteration

func Repeat(value string, times int) string {
	var acc string

	for i := 0; i < times; i++ {
		acc += value
	}

	return acc

}
