// Golang program to illustrate the usage of
// Sleep() function

// Including main package
package main

// Importing fmt and time
import (
	"fmt"
	"time"
)

// Main function
func main() {

	// Calling Sleep method
	time.Sleep(8 * time.Second)

	// Printed after sleep is over
	fmt.Println("Sleep Over.....")
}
