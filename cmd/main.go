package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// sleep function to enforce timestamp ordering
func sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

// getFilesAndDirs retrieves all files and directories in sorted order
func getFilesAndDirs(src string) ([]string, []string, error) {
	var files, dirs []string

	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, _ := filepath.Rel(src, path)
		if relPath == "." {
			return nil
		}
		if info.IsDir() {
			dirs = append(dirs, relPath)
		} else {
			files = append(files, relPath)
		}
		return nil
	})

	// Sort directories and files in natural order
	sort.Strings(dirs)
	sort.Strings(files)

	return files, dirs, err
}

// copyDirectories sequentially creates directories in the destination
func copyDirectories(src, dest string, dirs []string) error {
	for i, dir := range dirs {
		destPath := filepath.Join(dest, dir)
		err := os.MkdirAll(destPath, os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Printf("[DIR] %d/%d Created: %s\n", i+1, len(dirs), destPath)
		sleep(100) // Small delay for timestamp order
	}
	return nil
}

// copyFiles sequentially copies files
func copyFiles(src, dest string, files []string) error {
	for i, file := range files {
		srcPath := filepath.Join(src, file)
		destPath := filepath.Join(dest, file)

		// Ensure parent directories exist
		if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
			return err
		}

		// Copy file content
		input, err := os.ReadFile(srcPath)
		if err != nil {
			return err
		}
		if err := os.WriteFile(destPath, input, os.ModePerm); err != nil {
			return err
		}

		// Update timestamps
		now := time.Now()
		if err := os.Chtimes(destPath, now, now); err != nil {
			return err
		}

		fmt.Printf("[FILE] %d/%d Copied: %s\n", i+1, len(files), destPath)
		sleep(100) // Small delay for proper timestamp order
	}
	return nil
}

// promptInput asks the user for input with the given prompt and returns the input string
func promptInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// waitForExit waits for the user to press Enter before exiting
func waitForExit() {
	fmt.Print("\nPress Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	var src, dest string
	isInteractive := false

	// Check if arguments were provided
	if len(os.Args) >= 3 || (len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help")) {
		// Show usage if -h or --help flag is provided
		if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
			fmt.Println("Usage: qp <source> <destination>")
			os.Exit(0)
		}

		// Get source and destination from command line arguments
		src, dest = os.Args[1], os.Args[2]
	} else {
		isInteractive = true
		fmt.Println("🖥️  Interactive Mode")
		fmt.Println("--------------------")

		// Prompt for source path
		src = promptInput("Enter source path: ")
		if src == "" {
			fmt.Println("❌ Source path cannot be empty.")
			if isInteractive {
				waitForExit()
			}
			os.Exit(1)
		}

		// Prompt for destination path
		dest = promptInput("Enter destination path: ")
		if dest == "" {
			fmt.Println("❌ Destination path cannot be empty.")
			if isInteractive {
				waitForExit()
			}
			os.Exit(1)
		}
	}

	fmt.Printf("📂 Copying from: %s\n", src)
	fmt.Printf("📂 Copying to:   %s\n\n", dest)

	fmt.Println("🔍 Scanning source directory...")
	files, dirs, err := getFilesAndDirs(src)
	if err != nil {
		fmt.Println("Error scanning source:", err)
		if isInteractive {
			waitForExit()
		}
		os.Exit(1)
	}
	fmt.Printf("📁 Found %d directories and %d files.\n\n", len(dirs), len(files))

	fmt.Println("📂 Copying directories...")
	if err := copyDirectories(src, dest, dirs); err != nil {
		fmt.Println("Error copying directories:", err)
		if isInteractive {
			waitForExit()
		}
		os.Exit(1)
	}

	fmt.Println("\n📄 Copying files...")
	if err := copyFiles(src, dest, files); err != nil {
		fmt.Println("Error copying files:", err)
		if isInteractive {
			waitForExit()
		}
		os.Exit(1)
	}

	fmt.Println("\n✅ Copy completed successfully!")

	// If in interactive mode, wait for user to press Enter before exiting
	if isInteractive {
		waitForExit()
	}
}
