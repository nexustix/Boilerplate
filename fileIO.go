package boilerplate

import "os"

// FileExists tests if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

//CreateIfMissing creates a file if it does not exist
//deprecated
//func CreateIfMissing(filename string) {
//	CreateFileIfMissing(filename)
//}

//CreateFileIfMissing creates a file if it does not exist
func CreateFileIfMissing(filename string) {
	if !FileExists(filename) {
		outFile, err := os.Create(filename)
		FailError(err)
		outFile.Close()
	}
}
