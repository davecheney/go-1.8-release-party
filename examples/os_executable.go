package main

import (
	"fmt"
	"os"
)

func main() {
	// START OMIT
	fmt.Printf("os.Args[0]: %q\n", os.Args[0])
	exec, _ := os.Executable()
	fmt.Printf("os.Executable(): %q\n", exec)
	// END OMIT
}
