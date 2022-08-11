package main

import "golang.org/x/tour/pic"

func imageVal(x, y int) int {
	return x*x + y*y
}

func Pic(dx, dy int) [][]uint8 {
	var img [][]uint8
	for x := 0; x < dx; x++ {
		img = append(img, make([]uint8, dy))
		for y := 0; y < dy; y++ {
			img[x][y] = uint8(imageVal(x, y))
		}
	}
	return img
}

func main() {
	pic.Show(Pic)
}
