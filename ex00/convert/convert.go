// Package convert provides functionality for converting image files between formats.
//
// This package is designed to search for images with specific extensions and
// convert them to the target format.
package convert

import (
	"fmt"
	"image"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
)

// ImageConverter represents a tool for converting images between formats
type ImageConverter struct {
	// SourceExt is the file extension to search for (e.g., ".jpg")
	SourceExt string
	
	// TargetExt is the file extension to convert to (e.g., ".png")
	TargetExt string
}

// NewImageConverter creates a new ImageConverter instance
func NewImageConverter(sourceExt, targetExt string) *ImageConverter {
	return &ImageConverter{
		SourceExt: sourceExt,
		TargetExt: targetExt,
	}
}

// FindImageFiles searches for image files with the specified extension in the given directory
func (ic *ImageConverter) FindImageFiles(rootPath string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ic.SourceExt {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// GenerateOutputPath creates a new file path with the target extension
func (ic *ImageConverter) GenerateOutputPath(originalFilePath string) string {
	dir := filepath.Dir(originalFilePath)
	base := filepath.Base(originalFilePath)
	ext := filepath.Ext(base)
	fileNameWithoutExt := base[:len(base)-len(ext)]
	newFileName := fileNameWithoutExt + ic.TargetExt
	return filepath.Join(dir, newFileName)
}

// ConvertImage converts a single image file to the target format
func (ic *ImageConverter) ConvertImage(originalFilePath string) (string, error) {
	// Open the source file
	file, err := os.Open(originalFilePath)
	if err != nil {
		return "", fmt.Errorf("opening file %s: %w", originalFilePath, err)
	}
	defer file.Close()

	// Decode the image
	img, format, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("decoding image %s: %w", originalFilePath, err)
	}

	// Generate output path
	newFilePath := ic.GenerateOutputPath(originalFilePath)
	
	// Create the target file
	newFile, err := os.Create(newFilePath)
	if err != nil {
		return "", fmt.Errorf("creating file %s: %w", newFilePath, err)
	}
	defer newFile.Close()

	// For now, we only support PNG as output format
	err = png.Encode(newFile, img)
	if err != nil {
		return "", fmt.Errorf("encoding image to PNG: %w", err)
	}

	return fmt.Sprintf("Converted from %s format to PNG", format), nil
}

// ProcessImages converts all found image files to the target format
func (ic *ImageConverter) ProcessImages(rootPath string) (int, int, error) {
	files, err := ic.FindImageFiles(rootPath)
	if err != nil {
		return 0, 0, err
	}
	
	if len(files) == 0 {
		return 0, 0, fmt.Errorf("no %s files found in %s", ic.SourceExt, rootPath)
	}
	
	totalFiles := len(files)
	successCount := 0
	
	for _, filePath := range files {
		fmt.Printf("Converting %s...\n", filePath)
		result, err := ic.ConvertImage(filePath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("Success: %s -> %s (%s)\n", filePath, ic.GenerateOutputPath(filePath), result)
		successCount++
	}
	
	return successCount, totalFiles, nil
}
