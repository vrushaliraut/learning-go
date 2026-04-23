package main

import (
	"os"
)

func filelength_with_defer(filename string)(int64, error)  {

	// Open the file
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	// Schedule cleanup immediately
	defer f.Close()

	// Get file statistic to determine size
	// alternatively read all the bytes and count them - but start() is more efficient

	info, err := f.Stat()
	if err != nil  {
		return 0, err
	}

	// Return the size from fileInfo struct
	return info.Size(), nil
}