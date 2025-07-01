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

	pathToExamples := "../examples"

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
		fmt.Printf(fmt.Sprintf("ERROR: Example code must be 3 digits, got '%s'\n", arg))
		fmt.Printf(helpMessage1)
		fmt.Printf(helpMessage2)
		os.Exit(1)
	}

	// Find the Rust file
	pattern := fmt.Sprintf(pathToExamples + "/ex%s*.rs", arg)
	matches, err := filepath.Glob(pattern)
	if err != nil || len(matches) == 0 {
		fmt.Printf("> npm run rr %s\n\n", arg)
		fmt.Printf(fmt.Sprintf("ERROR: Example code not found in any file in examples\n"))
		fmt.Printf(helpMessage1)
		fmt.Printf(helpMessage2)
		os.Exit(1)
	}
	rustFile := matches[0]

	fmt.Printf("> npm run rr %s\n\n", arg)

	base := filepath.Base(pathToExamples + "/" + rustFile)
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
	lightGray := "\033[37;2m" // lighter gray (bright white, dimmed)
	yellow := "\033[33m"
	reset := "\033[0m"
	fmt.Println(lightGray + "====================================" + reset)
	name := base[:len(base)-len(filepath.Ext(base))]
	fmt.Println(lightGray + strings.ToUpper(name) + reset)
	fmt.Println(lightGray + "====================================" + reset)

	// Run the Rust binary and print its output in yellow
	cmdRun := exec.Command("./" + output)
	cmdRun.Stderr = os.Stderr

	// Capture stdout
	stdoutPipe, err := cmdRun.StdoutPipe()
	if err != nil {
		fmt.Println("Failed to capture output.")
		os.Exit(1)
	}
	if err := cmdRun.Start(); err != nil {
		fmt.Println("Execution failed.")
		os.Exit(1)
	}
	buf := make([]byte, 1024)
	for {
		n, err := stdoutPipe.Read(buf)
		if n > 0 {
			fmt.Print(yellow + string(buf[:n]) + reset)
		}
		if err != nil {
			break
		}
	}
	if err := cmdRun.Wait(); err != nil {
		fmt.Println("Execution failed.")
		os.Exit(1)
	}
	fmt.Println(lightGray + "====================================" + reset)
	fmt.Println()

	// Delete
	if err := os.Remove(output); err != nil {
		fmt.Println("Warning: could not delete binary:", err)
	}
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}
