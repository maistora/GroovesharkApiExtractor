package myIoUtil

import "os"

func AppendTextToFile(text, filename string) *os.File {
	if (!FileExists(filename)) {
		CreateFile(filename)
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
    	panic(err)
    }
    defer func() {
        if err := file.Close(); err != nil {
            panic(err)
        }
    }()

	if _, err := file.WriteString(text); err != nil {
	    panic(err)
	}
    
    return file
}

func CreateFile(filename string) {
	fo, err := os.Create(filename)
    if err != nil { panic(err) }
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
    }
    return true
}