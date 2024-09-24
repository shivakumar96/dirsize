package dirtests

import (
	"os"
	"testing"

	"dirsize.io/dirsize/display"
	dirread "dirsize.io/dirsize/read"
)

func initializeDir() (int, error) {
	err := os.MkdirAll("./dir1/dir2/dir3", 0755)
	if err != nil && !os.IsExist(err) {
		return 0, err
	}

	err = os.MkdirAll("./dir1/dir4", 0755)
	if err != nil && !os.IsExist(err) {
		return 0, err
	}

	f, err := os.OpenFile("./dir1/dir2/dir3/file1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var text = `
	This is a random text written to test the code. 
	This is a random text written to test the code. 
	This is a random text written to test the code. 
	`

	size1, err := f.Write([]byte(text))
	if err != nil {
		return 0, err
	}

	text = `
	This is a random text written to test the code. 
	This is a random text written to test the code. 
	This is a random text written to test the code. 
	This is a random text written to test the code.
	This  will used for testing.
	`

	f2, err := os.OpenFile("./dir1/dir4/lsfile2.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f2.Close()

	size2, err := f2.Write([]byte(text))
	if err != nil {
		return 0, err
	}
	return size1 + size2, nil

}

func cleandir() {
	os.RemoveAll("./dir1")
}
func TestReadDirectory(t *testing.T) {
	t.Cleanup(cleandir)
	size, err := initializeDir()
	if err != nil {
		t.Errorf("Cannot initilaize directory : %v", err.Error())
	}
	res, err := dirread.ReadDirectory("./dir1")
	if err != nil {
		t.Errorf("Cannot initilaize directory : %v", err.Error())
	}
	if res.TotalSize != int64(size) {
		display.PrintFormattedDirSizeResult([]*display.DirSizeResult{res}, 0, false, true)
		t.Errorf("directory size read incorrectly. expected %v but, got %v", size, res.TotalSize)
	}
}
