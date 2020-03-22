package libs

// Map generic function to populate an array with the results of the appender.
func Map(length int, appender func(int)) {
	for i := 0; i < length; i++ {
		appender(i)
	}
}
