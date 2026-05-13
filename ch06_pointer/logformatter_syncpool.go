package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

func main() {
	// simulate 5 log event
	for i := 100; i < 105; i++ {
		formatLog(i, "LOGIN_ATTEMPT")
	}
	// Notice in the output:
	// You'll likely see "Creating brand new buffer" only once or twice,
	// even though we "created" 5 logs. That's the pool working!
}

func formatLog(userID int, action string) {
	// 2. Borrow a buffer from the pool
	// If one exists, we get it. If not, New() is called.
	buf := logBufferPool.Get().(*bytes.Buffer)

	// 3. IMPORTANT: Always Reset the buffer!
	// Reused buffers still contain the "trash" from the last user.
	buf.Reset()

	// 4. Do our work
	buf.WriteString("TIME: ")
	buf.WriteString(time.Now().Format("15:04:05"))
	buf.WriteString(" | USER_ID: ")
	fmt.Fprintf(buf, "%d", userID)
	buf.WriteString(" | ACTION: ")
	buf.WriteString(action)

	// Print the result
	fmt.Println(buf.String())

	// 5. Return the buffer to the pool for the next goroutine to use
	logBufferPool.Put(buf)
}

// 1. Create the Pool
// We store *bytes.Buffer because it's a great tool for building strings
// without constant memory re-allocation.
var logBufferPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("✨ [Pool] Creating a brand new buffer...")
		return new(bytes.Buffer)
	},
}
