package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: runrust <file.rs>")
		os.Exit(1)
	}

	rustFile := os.Args[1]
	base := filepath.Base(rustFile)
	output := base[:len(base)-len(filepath.Ext(base))]

	// Compile
	cmdCompile := exec.Command("rustc", rustFile)
	cmdCompile.Stdout = os.Stdout
	cmdCompile.Stderr = os.Stderr
	if err := cmdCompile.Run(); err != nil {
		fmt.Println("Compilation failed.")
		os.Exit(1)
	}

	// Run
	cmdRun := exec.Command("./" + output)
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr
	if err := cmdRun.Run(); err != nil {
		fmt.Println("Execution failed.")
		os.Exit(1)
	}

	// Delete
	if err := os.Remove(output); err != nil {
		fmt.Println("Warning: could not delete binary:", err)
	}
}
