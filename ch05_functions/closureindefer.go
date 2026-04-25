package main

import (
	"fmt"
	"os"
)

func main() {
	writeFileSafely()
}

func writeFileSafely() {
	f, err := os.Create("important_data.txt")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}

	// Instead of just 'defer f.Close()', we defer a CLOSURE.
	defer func() {
		fmt.Println("Starting cleanup...")

		// Now we can capture and check the error from Close()
		closeErr := f.Close()
		if closeErr != nil {
			// We can log this critical failure!
			fmt.Printf("CRITICAL WARNING: Failed to close file safely: %v\n", closeErr)
		} else {
			fmt.Println("File closed successfully. Cleanup complete.")
		}
	}() // Don't forget the () to actually execute the closure!

	// Write some data to the file
	fmt.Fprintln(f, "Saving some very important data...")

	fmt.Println("Function logic finished, preparing to return...")
	// The deferred closure will run right after this line.
}
