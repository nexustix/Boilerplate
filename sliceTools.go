package boilerplate

// StringInSlice checks if string is element in given slice
func StringInSlice(theSlice []string, item string) bool {
	for _, v := range theSlice {
		if v == item {
			return true
		}
	}
	return false
}

// EliminateDuplicates eliminates duplicates in a given slice of strings
func EliminateDuplicates(theStrings []string) []string {
	var tmpStrings []string
	for _, v := range theStrings {
		if !StringInSlice(tmpStrings, v) {
			tmpStrings = append(tmpStrings, v)
		}
	}
	return tmpStrings
}

// EliminateStringInSlice eliminates all duplicates of a specific string
func EliminateStringInSlice(theStrings []string, deadString string) []string {
	var tmpStrings []string
	for _, v := range theStrings {
		if !(v == deadString) {
			tmpStrings = append(tmpStrings, v)
		}
	}
	return tmpStrings
}
