package libs

// Reduce generic function to reduce a slice on a given reducer function to a single output value of type int.
func Reduce(length int, reducer func(int), args ...int) {
	idx := 0
	if len(args) >= 1 {
		idx = args[0]
	}

	for i := idx; i < length; i++ {
		reducer(i)
	}
}
