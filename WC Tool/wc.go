package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wc() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())

	if len(input) == 1 {
		filename := input[0]
		fileInfo, statErr := os.Stat(filename)
		f, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			fmt.Fprintln(os.Stderr, "Error:", statErr)
		}
		if statErr != nil {
			fmt.Fprintln(os.Stderr, "Error:", statErr)
			return "Error!", filename
		}
		size := fileInfo.Size()
		lines := strings.Count(string(f), "\n")
		words := strings.Fields(string(f))
		return fmt.Sprintf("%d %d %d", lines, len(words), size), filename
	}

	if len(input) == 2 {
		toDo := input[0]
		filename := input[1]

		if toDo == "-c" {
			fileInfo, statErr := os.Stat(filename)
			if statErr != nil {
				fmt.Fprintln(os.Stderr, "Error:", statErr)
				return "Error!", filename
			}
			return fmt.Sprintf("%d", fileInfo.Size()), filename
		}

		f, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return "Error!", filename
		}

		if toDo == "-l" {
			lines := strings.Count(string(f), "\n")
			return fmt.Sprintf("%d", lines), filename
		}
		if toDo == "-w" {
			words := strings.Fields(string(f))
			return fmt.Sprintf("%d", len(words)), filename
		}
		if toDo == "-m" {
			chars := len([]rune(string(f)))
			return fmt.Sprintf("%d", chars), filename
		}
		return "Error!", filename
	}

	return "Error!", ""
}

func main() {
	result, filename := wc()
	fmt.Printf("%s %s\n", result, filename)
}
