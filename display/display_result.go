package display

import (
	"fmt"
	"strconv"
)

// Data structure to store result in the hierarchical manner
type DirSizeResult struct {
	DirName   string
	SubDirs   []*DirSizeResult
	TotalSize int64
}

func NewDirSizeResult(directoryname string, size int64) *DirSizeResult {
	dir := &DirSizeResult{DirName: directoryname, TotalSize: size,
		SubDirs: make([]*DirSizeResult, 0)}
	return dir
}

func (root *DirSizeResult) AppendSubDirResult(subDir *DirSizeResult) {
	root.SubDirs = append(root.SubDirs, subDir)
}

// used for formatted printing of results, stores the width
var (
	formatWidth = -1
	precision   = -1
)

// prints header
func PrintHeader(isHumanReadable bool) {
	width := formatWidth
	if isHumanReadable {
		width++
	}
	fmt.Printf("%*s %*s  directory name\n", width, "size", width, "cszie")
	fmt.Printf("-------------------------------\n")
}

// calculates the formatted width size, and is called only once
func UpdateDecimalWidth(directories []*DirSizeResult, isHumanReadable bool) {
	if formatWidth != -1 && precision != -1 {
		return
	}

	var cummulativeTotal int64 = 0
	for _, dir := range directories {
		cummulativeTotal += dir.TotalSize
	}

	formatWidth = len(strconv.Itoa(int(cummulativeTotal))) + 1
	if isHumanReadable {
		formatWidth = 6 // eg: max value can be 1023.9 => total width of 6. formatted right
		precision = 1
	} else {
		//value will be in bytes. The total size of all directories will be the maximum width
		formatWidth *= -1 // formatted to left
		precision = 0
	}
}

// prints formatted line of directory
func printFormattedLine(dir *DirSizeResult, cummulativeTotal int64, isHumanReadable bool) {
	//stores directory size
	var convertedVal float32
	var humanReadableSuffixVal string = ""
	var tempPrecision = precision

	// stores cummulative value
	var cummuConvertedVal float32
	var cummHumanReadableSuffixVal string = ""
	var cummTempPrecision = precision

	if isHumanReadable {
		convertedVal, humanReadableSuffixVal = ParseBytesToHumanReadable(dir.TotalSize)
		if humanReadableSuffixVal == "B" {
			tempPrecision = 0
		}
		cummuConvertedVal, cummHumanReadableSuffixVal = ParseBytesToHumanReadable(cummulativeTotal)
		if cummHumanReadableSuffixVal == "B" {
			cummTempPrecision = 0
		}

	} else {
		convertedVal = float32(dir.TotalSize)
		cummuConvertedVal = float32(cummulativeTotal)
	}
	// actual ouput
	sizestr := fmt.Sprintf("%*.*f%s %*.*f%s", formatWidth, tempPrecision, convertedVal, humanReadableSuffixVal,
		formatWidth, cummTempPrecision, cummuConvertedVal, cummHumanReadableSuffixVal)
	fmt.Printf("%s  %s\n", sizestr, dir.DirName)
}

// Prints the result of subdirectories recursively
func PrintFormattedDirSizeResult(directories []*DirSizeResult, cummulativeTotal int64, isHumanReadable bool, printRecurrsive bool) {

	// Base case: don't print if the list is empty
	if len(directories) == 0 {
		return
	}

	UpdateDecimalWidth(directories, isHumanReadable)
	for _, dir := range directories {
		if printRecurrsive {
			PrintFormattedDirSizeResult(dir.SubDirs, cummulativeTotal, isHumanReadable, printRecurrsive)
		}
		cummulativeTotal += dir.TotalSize
		printFormattedLine(dir, cummulativeTotal, isHumanReadable)
	}

}
