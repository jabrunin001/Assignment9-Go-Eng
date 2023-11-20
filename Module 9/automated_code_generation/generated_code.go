//go:generate go run github.com/example/codegenerator

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Automated Code Generation!")

	// Run the code generator
	err := runCodeGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Add your additional actions and functions here
}

func runCodeGenerator() error {
	// Specify the command to run the code generator
	cmd := exec.Command("codegenerator", "-input", "input.txt", "-output", "output.txt")

	// Set the working directory for the command
	cmd.Dir = "/path/to/codegenerator"

	// Set the environment variables for the command
	cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")

	// Run the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run code generator: %v, output: %s", err, output)
	}

	// Print the output of the code generator
	fmt.Println(string(output))

	return nil
}
