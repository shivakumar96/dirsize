package display

import (
	"fmt"
	"strconv"
)

// The data structure to store resut in the hierarchical manner
type DirSizeResult struct {
	DirName          string
	SubDirs          []DirSizeResult
	TotalSize        float32
	HumanReadableVal string // this will be empty if no human args is present
}

func NewDirSizeResult(directoryname string, size float32, humanReadbleVal string) *DirSizeResult {
	dir := &DirSizeResult{DirName: directoryname, TotalSize: size,
		SubDirs: make([]DirSizeResult, 0), HumanReadableVal: humanReadbleVal}
	return dir
}

func (root *DirSizeResult) AppendSubDirResult(subDir DirSizeResult) {
	root.SubDirs = append(root.SubDirs, subDir)
}

// used for formatted printing of results
var (
	formatWidth = -1
	precision   = -1
)

// This function will be called recursively for printing the result.
func PrintFormattedDirSizeResult(directories []DirSizeResult, isHumanReadable bool) {

	//base case doesn't print any if the directories slice lenth is empty
	if len(directories) == 0 {
		return
	}

	// calc
	var cummulativeTotal float32 = 0
	for _, dir := range directories {
		cummulativeTotal += dir.TotalSize
	}

	// this will be initialized only onve at the beginning
	if formatWidth == -1 && precision == -1 {
		formatWidth = len(strconv.Itoa(int(cummulativeTotal))) + 1
		if isHumanReadable {
			formatWidth = 6 // max value can be 1023.9 total width of 6. formatted right
			precision = 1
		} else {
			// everything will be in bytes total size of all directories will be the max width
			formatWidth *= -1 //foramted to left
			precision = 0
		}
	}

	//print recursively all the subdirectories value

	for _, dir := range directories {
		sizestr := fmt.Sprintf("%*.*f%s", formatWidth, precision, dir.TotalSize, dir.HumanReadableVal)
		fmt.Printf("%s %s", sizestr, dir.DirName)
		PrintFormattedDirSizeResult(dir.SubDirs, isHumanReadable)

	}

	//convert

	//fmt.printf("%*.*f%s Total", formatWidth, precision, cummulativeTotal, directories[])

}
