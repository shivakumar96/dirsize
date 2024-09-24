package main

/**
Author : Shivakumar Suresh
**/

import (
	"flag"
	"fmt"

	"dirsize.io/dirsize/display"
	dirread "dirsize.io/dirsize/read"
)

func printHelp() {
	var help = `
	Name:
	dirsize - calculates the directory size and prints it
	   
	Usage:
	dirsize [Options] <dirirectory list>
	   
	Options:
	--human      Prints size in human-readable format
	--recursive  Prints size of subdirectories recursively
	--help       Help command

	eg: dirsize --human --recursive dir1 dir2

	Output: 
	size  cszie  directory name
	-------------------------------
	
	size           - the size of the file
	cszie          - cummulative size 
	directory name - the name of the directory
	`
	fmt.Println(help)
}

// Entry point for the command line utility, parses the input and invokes the functions
func main() {
	isHumanReadable := flag.Bool("human", false, "Prints size in human readable format")
	isRecurssive := flag.Bool("recursive", false, "Prints size of subdirectories recursively")

	flag.Parse()
	directories := flag.Args()

	if len(directories) == 0 {
		printHelp()
		return
	}

	dirres := dirread.ReadAllDirectory(directories)
	if len(dirres) == 0 {
		return
	}

	display.UpdateDecimalWidth(dirres, *isHumanReadable)
	display.PrintHeader(*isHumanReadable)
	display.PrintFormattedDirSizeResult(dirres, 0, *isHumanReadable, *isRecurssive)

}
