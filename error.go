package boilerplate

// GotError true if an Error occured
func GotError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

// FailError panics if an Error occured
func FailError(err error) {
	if err != nil {
		//panic("Error")
		panic(err)
	}
}
