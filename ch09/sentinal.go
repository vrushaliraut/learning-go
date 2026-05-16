package main

import (
	"bytes"
	"errors"
	"fmt"
)

var (
	// ErrInvalidHeader indicates the file doesn't have a valid ZIP signature
	ErrInvalidHeader = errors.New("demo/zip: invalid zip header signature")

	// ErrEmptyArchiver indicates the reader encountered zero-length data
	ErrEmptyArchive = errors.New("demo/zip: empty archive")

	// ErrCorrupted indicates the central directory is malformed
	ErrCorrupted = errors.New("demo/zip: corrupted central directory")
)

type SimpleZipReader struct {
	data []byte
}

// NewSimpleZipReader creates a new reader or returns a sentinel error
func NewSimpleZipReader(data []byte) (*SimpleZipReader, error) {
	// Check 1: Empty data
	if len(data) == 0 {
		// Return the EXACT package-level variable (not a new error!)
		return nil, ErrEmptyArchive
	}

	// Check 2: Valid ZIP signature (PK at offset 0)
	if len(data) < 4 || !bytes.HasPrefix(data, []byte("PK\x03\x04")) {
		// Return the EXACT package-level variable
		return nil, ErrInvalidHeader
	}

	// Check 3: Minimum viable structure (local file header + file name + data)
	if len(data) < 30 {
		return nil, ErrCorrupted
	}

	return &SimpleZipReader{data: data}, nil
}

func processZipFile(data []byte) error {
	reader, err := NewSimpleZipReader(data)

	if err == ErrEmptyArchive {
		fmt.Println("File is empty. Please provide a non-empty file.")
		return err
	}

	if err == ErrInvalidHeader {
		fmt.Println("Not a zip file please provide a valid .zip archive")
		return err
	}

	if err == ErrCorrupted {
		fmt.Println("File is corrupted. Please provide a valid .zip archive")
		return err
	}

	if err != nil {
		fmt.Errorf("unexpected zip error: %w", err)

	}
	fmt.Printf("Success: valid ZIP structure detected (%d bytes)\n", len(reader.data))
	return nil
}

func main() {
	fmt.Println("1: Empty file")
	processZipFile([]byte{})

	fmt.Println("2: Invalid Header")
	processZipFile([]byte("This looks like a text not a zip file "))

	fmt.Println("3: Corrupted")
	processZipFile([]byte("PK\x03\x04"))

	fmt.Println("5: Valid Zip File")
	validData := make([]byte, 50)
	copy(validData, []byte("PK\x03\x04"))

	processZipFile(validData)
}
