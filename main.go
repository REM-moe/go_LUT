package main

import (
	"fmt"
	"os"
)

type Lut3D struct {
	Title string
	Size  int
	Data  [][][]Color
}

type Color struct {
	R, G, B float64
}

func main() {

	// os.Args just like java where first element is the name of the program
	// and the rest are the arguments passed to the program

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <path_to_images_folder> <path_to_lut_file>")
		return
	}

	imagesDir := os.Args[1]
	lutFile := os.Args[2]

	fmt.Printf("**************** \n FOLDER : %s  \n**************** \n LUT FILE : %s", imagesDir, lutFile)
}

func processImage(image string) (*Lut3D, error) {
	return nil, nil
}
