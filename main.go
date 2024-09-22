package main

import (
	"fmt"

	"dirsize.io/dirszie/display"
)

// Enrty point for the command line utility
// parses the input and invoke the functions
func main() {
	s, h := display.ParseBytesToHumanReadable(42949672960)
	fmt.Printf("%*.*f %v\n", 5, 2, s, h)
	fmt.Printf("%v ", display.ParseHumanReadableToBytes(s, h))
}
