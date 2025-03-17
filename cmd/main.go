package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// Sleep function to enforce timestamp ordering
func sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

// Get all files and directories in sorted order
func getFilesAndDirs(src string) ([]string, []string, error) {
	var files []string
	var dirs []string

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
	sort.Slice(dirs, func(i, j int) bool {
		return strings.Compare(dirs[i], dirs[j]) < 0
	})
	sort.Slice(files, func(i, j int) bool {
		return strings.Compare(files[i], files[j]) < 0
	})

	return files, dirs, err
}

// Copy directories sequentially
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

// Copy files sequentially
func copyFiles(src, dest string, files []string) error {
	for i, file := range files {
		srcPath := filepath.Join(src, file)
		destPath := filepath.Join(dest, file)

		// Ensure parent directories exist
		err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
		if err != nil {
			return err
		}

		// Copy file
		input, err := os.ReadFile(srcPath)
		if err != nil {
			return err
		}
		err = os.WriteFile(destPath, input, os.ModePerm)
		if err != nil {
			return err
		}

		// Update timestamps
		now := time.Now()
		err = os.Chtimes(destPath, now, now)
		if err != nil {
			return err
		}

		fmt.Printf("[FILE] %d/%d Copied: %s\n", i+1, len(files), destPath)
		sleep(100) // Small delay for timestamp order
	}
	return nil
}

func main() {
	// Get CLI arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: qp <source> <destination>")
		os.Exit(1)
	}
	src := os.Args[1]
	dest := os.Args[2]

	fmt.Printf("📂 Copying from: %s\n", src)
	fmt.Printf("📂 Copying to:   %s\n\n", dest)

	// Get files and directories
	fmt.Println("🔍 Scanning source directory...")
	files, dirs, err := getFilesAndDirs(src)
	if err != nil {
		fmt.Println("Error scanning source:", err)
		os.Exit(1)
	}
	fmt.Printf("📁 Found %d directories and %d files.\n\n", len(dirs), len(files))

	// Copy directories
	fmt.Println("📂 Copying directories...")
	err = copyDirectories(src, dest, dirs)
	if err != nil {
		fmt.Println("Error copying directories:", err)
		os.Exit(1)
	}

	// Copy files
	fmt.Println("\n📄 Copying files...")
	err = copyFiles(src, dest, files)
	if err != nil {
		fmt.Println("Error copying files:", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ Copy completed successfully!")
}
