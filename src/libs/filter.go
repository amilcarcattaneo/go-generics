package libs

// Filter generic function to filter a slice on a given predicate.
func Filter(length int, predicate func(int) bool, appender func(int)) {
	for i := 0; i < length; i++ {
		if predicate(i) {
			appender(i)
		}
	}
}
