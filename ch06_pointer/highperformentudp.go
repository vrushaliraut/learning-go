package main

/*
In high-throughput Go applications, the Garbage Collector (GC) is your biggest enemy.
If you allocate a new buffer 100,000 times per second,
the GC will spend all its time "sweeping" those short-lived objects,
leading to "Stop the World" pauses that kill your application's responsiveness.

By reusing a single buffer, you keep the memory on the stack (or a single stable heap allocation),
resulting in 0% allocation overhead inside your hot loop.
*/

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Setup UDP address and connection

	addr, _ := net.ResolveUDPAddr("udp", ":5140")
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Printf("Error listening on UDP: %s\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("Syslog receiver started on :5140")

	// 2. PRE-ALLOCATION: The "Golden Rule"
	// We allocate once outside the loop.
	// 65535 is the max size of a UDP packet.

	buffer := make([]byte, 65535)

	for {
		// 3. MUTATION: ReadFromUDP writes directly into our pre-allocated memory.
		// n is the number of bytes actually read.

		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading from UDP: %s\n", err)
			continue
		}

		// 4. WINDOWING: Pass a sub-slice (view) of the buffer.
		// This does NOT copy the data. It just creates a new header
		// pointing to the same backing array.
		processLog(remoteAddr, buffer[:n])
	}
}

func processLog(addr *net.UDPAddr, data []byte) {
	// In a real app, you'd parse the syslog format here.
	// We use string(data) here for the demo, but in production,
	// even this string conversion might be an allocation you'd want to avoid!
	fmt.Printf("[%s] Received: %s\n", addr, string(data))
}
