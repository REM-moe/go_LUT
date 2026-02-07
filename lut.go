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

func loadLUT(filename string) (*Lut3D, error) {

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lut := &Lut3D{}

	headerPassed := false

	for scanner.Scan() {

		line := scanner.Text()
		words := strings.Fields(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.Contains(line, "LUT_3D_SIZE") && len(words) >= 2 && !headerPassed {

			text, err := strconv.Atoi(words[1])

			if err != nil {
				return nil, err
			}

			lut.Size = text
			lut.Data = make([]Color, 0, text*text*text)
			headerPassed = true

		} else if len(words) == 3 && headerPassed {
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
