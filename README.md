[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/tsdoc)](https://goreportcard.com/report/github.com/Konstantin8105/tsdoc)
[![GoDoc](https://godoc.org/github.com/Konstantin8105/tsdoc?status.svg)](https://godoc.org/github.com/Konstantin8105/tsdoc)
![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)

 # Triplet-splash

 Get documentation from Go source

 ## Installation

```
 go get -u github.com/Konstantin8105/tsdoc
 cd $GOPATH/Konstantin8105/tsdoc/cmd
 go install
```

 Example Go source with triplet-slash:
```go
func add(a, b int) (int, error) {
	/// Function `add` return summ of two positive integer values.
	///
	if a < 0 || b < 0 {
		/// If some of value is negative, then return the error.
		return -1, fmt.Errorf("Some value is negative")
	}
	return a + b, nil
}
```

Output:

```
 Function `add` return summ of two positive integer values.

 If some of value is negative, then return the error.
```

 ## Example of generation README.md on linux

 Insert into your Go code:

```go
//go:generate sh -c "tsdoc > README.md"
```

 ## Function Get
 Function Get search all Go files in `path` and go deeper by folders.

 If cannot find absolute path, then return error.

 If `path` is not exist, then return error.


 If `path` is not the folder, then return error.


 ## Searching.


 List of ignore folders: vendor, .git


 Searching run from folder `path`.


 For avoid infinite loop added limits of search iterations(cycles).


 If cannot read directory, then return error.


 Searching only Go files.


 If cannot find any acceptable files, then return error.


 ## Sorting.

 Before reading all files, start a sorting of filename.

 For example: at the begin read a file with name `complex.go`,
 then read file `complex_test.go`.


 ## Read all files.

 Reading files one by one.


 If cannot read a file content, then return the error.


 Read file line by line.


 Before triplet-slash is not acceptable any characters,
 except `\t` or space.


