package main

import (
	"image"
	"image/color"
)

func applyLUT(src image.Image, lut *Lut3D) *image.RGBA {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	lutSize := lut.Size
	maxIndex := float64(lutSize - 1)
	planeSize := lutSize * lutSize // Pre-calculate stride for Blue

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			c := src.At(x, y)
			r, g, b, a := c.RGBA()

			// we normalize each pixel
			rF := float64(r) / 65535.0
			gF := float64(g) / 65535.0
			bF := float64(b) / 65535.0

			// get the 3d cordinates
			idxR := int(rF * maxIndex)
			idxG := int(gF * maxIndex)
			idxB := int(bF * maxIndex)

			// Flat Index Math: (B * Area) + (G * Width) + R
			// In LUT files the first column is red , second green, thrid blue
			// so when we converted them to a 1d arrat [red, green, blue]
			index := (idxB * planeSize) + (idxG * lutSize) + idxR

			// Bounds check (just in case of float precision errors)
			if index < 0 {
				index = 0
			}
			if index >= len(lut.Data) {
				index = len(lut.Data) - 1
			}

			newColor := lut.Data[index]

			dst.Set(x, y, color.RGBA{
				R: uint8(newColor.R * 255),
				G: uint8(newColor.G * 255),
				B: uint8(newColor.B * 255),
				A: uint8(a >> 8), // Keep original Alpha
			})
		}
	}
	return dst
}
