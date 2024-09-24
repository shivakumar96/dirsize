package dirtests

import (
	"testing"

	"dirsize.io/dirsize/display"
)

func TestParseBytesToHumanReadable(t *testing.T) {

	// conversion to KB
	var byteVal int64 = 4096
	ans, suffix := display.ParseBytesToHumanReadable(byteVal)
	if ans != 4 || suffix != "K" {
		t.Errorf("Incorrect convertion of bytes to human readable value of  %v to %v %v\n", byteVal, ans, suffix)
	}

	// conversion to MB
	byteVal = 6291456
	ans, suffix = display.ParseBytesToHumanReadable(byteVal)
	if ans != 6 || suffix != "M" {
		t.Errorf("Incorrect convertion of bytes to human readable value of  %v to %v %v\n", byteVal, ans, suffix)
	}

	// conversion to GB
	byteVal = 8589934592
	ans, suffix = display.ParseBytesToHumanReadable(byteVal)
	if ans != 8 || suffix != "G" {
		t.Errorf("Incorrect convertion of bytes to human readable value of  %v to %v %v\n", byteVal, ans, suffix)
	}
}
