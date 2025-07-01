package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	clearTerminal()

	helpMessage1 := "USAGE: send 3-digit example code\n"
	helpMessage2 := "EXAMPLE: npm run rr 040\n\n"
	if len(os.Args) != 2 {
		fmt.Printf("> npm run rr\n\n")
		fmt.Printf(helpMessage1)
		fmt.Printf(helpMessage2)
		os.Exit(1)
	}

	arg := os.Args[1]
	re := regexp.MustCompile(`^\d{3}$`)
	if !re.MatchString(arg) {
		fmt.Printf("> npm run rr %s\n\n", arg)
		fmt.Printf(fmt.Sprintf("ERROR: Example code not correct\n"))
		fmt.Printf(helpMessage1)
		fmt.Printf(helpMessage2)
		os.Exit(1)
	}

	// Find the Rust file
	pattern := fmt.Sprintf("ex%s*.rs", arg)
	matches, err := filepath.Glob(pattern)
	if err != nil || len(matches) == 0 {
		fmt.Printf("> npm run rr %s\n\n", arg)
		fmt.Printf("ERROR: No file found matching pattern %s\n", pattern)
		fmt.Printf(helpMessage1)
		fmt.Printf(helpMessage2)
		os.Exit(1)
	}
	rustFile := matches[0]

	fmt.Printf("> npm run rr %s\n\n", arg)

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
	fmt.Println("====================================")
	name := base[:len(base)-len(filepath.Ext(base))]
	fmt.Println(strings.ToUpper(name))
	fmt.Println("====================================")
	cmdRun := exec.Command("./" + output)
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr
	if err := cmdRun.Run(); err != nil {
		fmt.Println("Execution failed.")
		os.Exit(1)
	}
	fmt.Println("====================================")

	// Delete
	if err := os.Remove(output); err != nil {
		fmt.Println("Warning: could not delete binary:", err)
	}
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}
