package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func sizeToBytes(sizeStr string) (int64, error) {
	re := regexp.MustCompile(`^(\d+)\s*(B|KB|MB|GB)$`)
	matches := re.FindStringSubmatch(strings.TrimSpace(sizeStr))
	if matches == nil {
		return 0, fmt.Errorf("invalid size format. Expected format: NUMBER[B|KB|MB|GB] (e.g., 10KB)")
	}

	value, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %v", err)
	}

	switch matches[2] {
	case "B":
		return value, nil
	case "KB":
		return value * 1024, nil
	case "MB":
		return value * 1024 * 1024, nil
	case "GB":
		return value * 1024 * 1024 * 1024, nil
	default:
		return 0, fmt.Errorf("unsupported unit: %s", matches[2])
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <size> <output-file>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s 10MB output.bin\n", os.Args[0])
		os.Exit(1)
	}

	size, err := sizeToBytes(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing size: %v\n", err)
		os.Exit(1)
	}

	outputPath := os.Args[2]
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	const bufferSize = 1024 * 1024
	buffer := make([]byte, bufferSize)
	remaining := size

	for remaining > 0 {
		writeSize := bufferSize
		if remaining < bufferSize {
			writeSize = int(remaining)
		}

		rand.Read(buffer[:writeSize])

		n, err := file.Write(buffer[:writeSize])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
			os.Exit(1)
		}
		remaining -= int64(n)
	}

	fmt.Printf("Successfully created random binary file: %s (%d bytes)\n", outputPath, size)
}
