// Package tsdoc get documentation from Go source with triple-slash.
package tsdoc

import (
	"bytes"
	"fmt"
	"github.com/Konstantin8105/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

//go:generate sh -c "go run cmd/main.go > README.md"

///[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/tsdoc)](https://goreportcard.com/report/github.com/Konstantin8105/tsdoc)
///[![GoDoc](https://godoc.org/github.com/Konstantin8105/tsdoc?status.svg)](https://godoc.org/github.com/Konstantin8105/tsdoc)
///![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)
///
/// # Triplet-splash
///
/// Get documentation from Go source
///
/// ## Installation
///
///```
/// go get -u github.com/Konstantin8105/tsdoc
/// cd $GOPATH/Konstantin8105/tsdoc/cmd
/// go install
///```
///
/// Example Go source with triplet-slash:
///```go
///func add(a, b int) (int, error) {
///	/// Function `add` return summ of two positive integer values.
///	///
///	if a < 0 || b < 0 {
///		/// If some of value is negative, then return the error.
///		return -1, fmt.Errorf("Some value is negative")
///	}
///	return a + b, nil
///}
///```
///
///Output:
///
///```
/// Function `add` return summ of two positive integer values.
///
/// If some of value is negative, then return the error.
///```
///
/// ## Example of generation README.md on linux
///
/// Insert into your Go code:
///
///```go
/////go:generate sh -c "tsdoc > README.md"
///```

//go:generate sh -c "go run cmd/main.go > README.md"

var separator string = string(filepath.Separator)

// Get return documentation from Go source with triple-slash.
//
// Variable `path` is folder name with Go source code.
// Variable `deep`. If true, then search triplet-slash documentation in internal
// foladers.
//
// For example:
//
//		func add(a, b int) (int, error) {
//			/// Function `add` return summ of two positive integer values.
//			if a < 0 || b < 0 {
//				/// If some of value is negative, then return the error.
//				return -1, fmt.Errorf("Some value is negative")
//			}
//			return a + b, nil
//		}
//
// Return:
//
//		Function `add` return summ of two positive integer values.
//		If some of value is negative, then return the error.
//
func Get(path string, deep bool) (doc string, err error) {
	defer func() {
		if err != nil {
			et := errors.New("TSDOC")
			et.Add(fmt.Errorf("Error with input data: %v %v", path, deep))
			et.Add(err)
			err = et
		}
	}()
	///
	/// ## Function Get
	/// Function Get search all Go files in `path` and go deeper by folders.
	///
	{
		var apath string
		apath, err = filepath.Abs(path)
		if err != nil {
			/// If cannot find absolute path, then return error.
			return "", fmt.Errorf("Cannot get absolute path: %v", err)
		}
		var st os.FileInfo
		st, err = os.Stat(apath)
		if os.IsNotExist(err) {
			///
			/// If `path` is not exist, then return error.
			///
			return "", fmt.Errorf("Cannot find: `%s`", path)
		}
		if !st.IsDir() {
			///
			/// If `path` is not the folder, then return error.
			///
			return "", fmt.Errorf("Is not a folder: `%s`", path)
		}
	}
	///
	/// ## Searching.
	///
	var files []string
	{
		///
		/// List of ignore folders: vendor, .git
		///
		ignore := []string{"vendor", ".git"}
		///
		/// Searching run from folder `path`.
		///
		folderList := []string{path}
		///
		/// For avoid infinite loop added limits of search iterations(cycles).
		///
		for iter := 0; iter < 1000000; iter++ {
			findFolders := []string{}
			for _, folder := range folderList {
				fileInfo, err := ioutil.ReadDir(folder)
				if err != nil {
					///
					/// If cannot read directory, then return error.
					///
					return "", fmt.Errorf("Cannot read dir `%s`: %v", folder, err)
				}
				for _, f := range fileInfo {
					if f.IsDir() {
						isIgnore := false
						for _, ig := range ignore {
							if f.Name() == ig {
								isIgnore = true
								break
							}
						}
						if isIgnore {
							continue
						}
						findFolders = append(findFolders, f.Name())
						continue
					}
					if name := f.Name(); strings.HasSuffix(name, ".go") {
						///
						/// Searching only Go files.
						///
						files = append(files, folder+separator+name)
					}
				}
			}
			folderList = findFolders
			if !deep || len(folderList) == 0 {
				break
			}
		}
	}

	if len(files) == 0 {
		///
		/// If cannot find any acceptable files, then return error.
		///
		return "", fmt.Errorf("Cannot find any files")
	}

	///
	/// ## Sorting.
	///
	/// Before reading all files, start a sorting of filename.
	///
	/// For example: at the begin read a file with name `complex.go`,
	/// then read file `complex_test.go`.
	///
	sort.Strings(files)

	///
	/// ## Read all files.
	///
	/// Reading files one by one.
	///
	for _, filename := range files {
		var content []byte
		content, err = ioutil.ReadFile(filename)
		if err != nil {
			///
			/// If cannot read a file content, then return the error.
			///
			return "", fmt.Errorf("Cannot read file content: %v", filename)
		}
		///
		/// Read file line by line.
		///
		lines := bytes.Split(content, []byte("\n"))
		for i := range lines {
			line := lines[i]
			const ts string = "///"
			index := bytes.Index(line, []byte(ts))
			if index < 0 {
				// that line haven`t triplet-slash
				continue
			}
			if index > 0 {
				///
				/// Before triplet-slash is not acceptable any characters,
				/// except `\t` or space.
				///
				isAcceptableLine := true
				for pos := 0; pos < index; pos++ {
					if !(line[pos] == ' ' || line[pos] == '\t') {
						isAcceptableLine = false
						break
					}
				}
				if !isAcceptableLine {
					continue
				}
			}
			doc += string(line[index+len(ts):]) + "\n"
		}
	}
	return
}
