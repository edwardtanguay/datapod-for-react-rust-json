package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func main() {
	clearTerminal()

	helpMessage := "Usage: npm run rr 040\n\n"
	if len(os.Args) != 2 {
		fmt.Printf(helpMessage)
		os.Exit(1)
	}

	arg := os.Args[1]
	re := regexp.MustCompile(`^\d{3}$`)
	if !re.MatchString(arg) {
		fmt.Printf(helpMessage)
		os.Exit(1)
	}

	rustFile := arg
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

func clearTerminal() {
    fmt.Print("\033[H\033[2J")
}