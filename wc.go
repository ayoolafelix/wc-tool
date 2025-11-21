package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func wc() (string, string) {
	args := os.Args[1:]
	wantM, wantL, wantC, wantW := false, false, false, false
	var filename string

	for _, a := range args {
		switch a {
		case "-m":
			wantM = true
		case "-l":
			wantL = true
		case "-c":
			wantC = true
		case "-w":
			wantW = true
		default:
			if filename == "" {
				filename = a
			}

		}
	}

	if !wantL && !wantC && !wantW && !wantM {
		wantL = true
		wantC = true
		wantW = true
	}

	var data []byte
	var err error
	sourceIsFile := false

	//if there is file name
	if filename != "" {
		data, err = os.ReadFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading file ", err)
			return "Error", filename
		}
		sourceIsFile = true
	} else { //no file name, read from stdin
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading stdin: ", err)
			return "Error", ""
		}
	}

	bytesCount := len(data)
	linesCount := strings.Count(string(data), "\n")
	if len(data) > 0 && data[len(data)-1] != '\n' {
		linesCount++
	}
	wordsCount := len(strings.Fields(string(data)))
	charsCount := len([]rune(string(data)))

	var output []string
	if wantL {
		output = append(output, fmt.Sprintf("%d", linesCount))
	}
	if wantW {
		output = append(output, fmt.Sprintf("%d", wordsCount))
	}
	if wantC {
		output = append(output, fmt.Sprintf("%d", bytesCount))
	}
	if wantM {
		output = append(output, fmt.Sprintf("%d", charsCount))
	}

	result := strings.Join(output, " ")

	if sourceIsFile {
		return result, filename
	} else {
		return result, ""
	}

}

func main() {
	result, filename := wc()
	if filename == "" {
		fmt.Printf("%7s\n", result)
	} else {
		fmt.Printf("%7s %s\n", result, filename)
	}

}
