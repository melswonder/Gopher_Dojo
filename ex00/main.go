// Package main provides the entry point for the image conversion tool.
//
// This program searches for image files with a specific extension
// and converts them to another format.
package main

import (
	"fmt"
	_ "image/jpeg" // Register JPEG decoder
	"os"

	"github.com/melswonder/Gopher_Dojo/ex00/convert" // 自作パッケージ
)

func main() {
	path := getValidDirName()
	
	// 自作パッケージのユーザー定義型を使用
	converter := convert.NewImageConverter(".jpg", ".png")
	
	// 変換処理の実行
	successCount, totalFiles, err := converter.ProcessImages(path)
	
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("\nConversion complete: %d of %d files successfully converted to PNG.\n", 
		successCount, totalFiles)
}

// getValidDirName validates and returns the directory path from command line arguments
func getValidDirName() string {
	if len(os.Args) != 2 {
		fmt.Println("Error: invalid argument")
		fmt.Println("Usage: go run main.go <directory_path>")
		os.Exit(1)
	}
	return os.Args[1]
}