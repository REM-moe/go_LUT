package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <images_dir> <lut_path>")
		return
	}

	imgDir := os.Args[1]
	lutPath := os.Args[2]

	// 1. Load LUT (Once)
	fmt.Println("Loading LUT...")
	lut, err := loadLUT(lutPath)
	if err != nil {
		panic(err)
	}

	// 2. Read Directory
	files, err := os.ReadDir(imgDir)
	if err != nil {
		panic(err)
	}

	// 3. Create Output Directory
	outDir := "processed"
	os.Mkdir(outDir, 0755)

	// 4. Concurrency Setup
	var wg sync.WaitGroup

	fmt.Println("Processing images...")

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Filter for images
		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			continue
		}

		// Increment counter
		wg.Add(1)

		go func(filename string) {
			defer wg.Done() // Decrement counter when finished

			inPath := filepath.Join(imgDir, filename)
			f, err := os.Open(inPath)
			if err != nil {
				fmt.Printf("Failed to open %s: %v\n", filename, err)
				return
			}
			defer f.Close()

			src, _, err := image.Decode(f)
			if err != nil {
				return
			}

			dst := applyLUT(src, lut)

			outPath := filepath.Join(outDir, filename+".png")
			outF, err := os.Create(outPath)
			if err != nil {
				fmt.Printf("Failed to create %s: %v\n", outPath, err)
				return
			}
			defer outF.Close()

			png.Encode(outF, dst)
			fmt.Printf("Processed: %s\n", filename)
		}(file.Name())
	}

	wg.Wait()
	fmt.Println("All done!")
}
