package dirread

import (
	"os"

	"path/filepath"

	"dirsize.io/dirsize/display"

	dErr "dirsize.io/dirsize/errors"
)

/**

NOTE: The other way of doing it is to read the directory without using the walk function but recursively,
	 and calculate the size directly print it, and go to the following directory

	 The idea of storing the result here is that we could use it to change the print format and easily traverse
	 the tree structure of the directory.


**/

// calculates the final size, including the sub-directory size
func calculateTotalSize(dir *display.DirSizeResult) int64 {

	if dir == nil {
		return 0
	}

	var totalChildMem int64 = 0

	for _, child := range dir.SubDirs {
		totalChildMem += calculateTotalSize(child)
	}
	dir.TotalSize += totalChildMem

	return dir.TotalSize
}

// The reads a directory and its content and find the directory size by adding file size
func ReadDirectory(directoryPath string) (*display.DirSizeResult, error) {

	var root = directoryPath

	// stores the refernces
	diMap := make(map[string]*display.DirSizeResult)
	err := filepath.Walk(root, func(path string, fInfo os.FileInfo, err error) error {

		if err != nil {
			return dErr.ErrDirCannotRead
		}
		parentPath := filepath.Dir(path)
		prev, found := diMap[parentPath]

		//If it is a file, read the size of the file
		//(parent DirSizeResult reference will be present in the map, it just updates it)
		if !fInfo.IsDir() {
			prev.TotalSize += fInfo.Size()
		}

		//If it is a directory, so create a new DirSizeResult referring to the directory and add it to the map
		if fInfo.IsDir() {
			//create new refernce to DirSizeResult
			curr := display.NewDirSizeResult(path, 0)

			// valides conditions if root is ./ ../ "././../." and updates the root accordingly
			if fInfo.Name() == "." || fInfo.Name() == ".." {
				if parentPath == ".." && fInfo.Name() == "." {
					diMap[parentPath] = curr
					root = parentPath
				} else {
					diMap[fInfo.Name()] = curr
					root = fInfo.Name()
				}

			} else {
				diMap[path] = curr
			}
			// If parent is present append current reference to parent sub directory
			if found {
				prev.AppendSubDirResult(curr)

			}
		}

		return nil
	})

	if err != nil {
		return nil, dErr.ErrDirCannotRead
	}
	//calculates the total size of the directory
	calculateTotalSize(diMap[root])
	return diMap[root], nil
}

// reads slice of directories
func ReadAllDirectory(dirs []string) []*display.DirSizeResult {
	result := make([]*display.DirSizeResult, 0)
	for _, dir := range dirs {
		dirRes, err := ReadDirectory(dir)
		if err != nil {
			dErr.PrintError(dir, err)
			continue
		}
		result = append(result, dirRes)
	}
	return result
}
