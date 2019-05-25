 # Triplet-splash

 Get documentation from Go source

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

