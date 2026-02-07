package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <path_to_images_folder> <path_to_lut_file>")
		return
	}

	imagesDir := os.Args[1]
	lutFile := os.Args[2]

	fmt.Printf(" FOLDER : %s \n LUT FILE : %s\n", imagesDir, lutFile)

	_, err := loadLUT(lutFile)
	if err != nil {
		fmt.Print(err)
		return
	}

}
