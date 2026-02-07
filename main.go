package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lut3D struct {
	Title string
	Size  int
	Data  []Color
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

	fmt.Printf(" FOLDER : %s \n LUT FILE : %s\n", imagesDir, lutFile)

	loadLUT(lutFile)
}

func loadLUT(filename string) (*Lut3D, error) {

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lut := &Lut3D{}

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		// check word starts with LUT_3D_SIZE
		if strings.Contains(line, "LUT_3D_SIZE") && len(words) >= 2 {

			text, err := strconv.Atoi(words[1])

			if err != nil {
				return nil, err
			}

			lut.Size = text
			lut.Data = make([]Color, 0, text*text*text)

		} else if len(words) == 3 {
			r, err1 := strconv.ParseFloat(words[0], 64)
			g, err2 := strconv.ParseFloat(words[1], 64)
			b, err3 := strconv.ParseFloat(words[2], 64)

			if err1 != nil || err2 != nil || err3 != nil {
				return nil, fmt.Errorf("bad color data: %s", line)
			}

			lut.Data = append(lut.Data, Color{
				R: r,
				G: g,
				B: b,
			})

		} else {
			continue
		}
	}

	return lut, nil

}
