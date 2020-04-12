package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for y := range ret {
		for x := 0; x < dy; x++ {
			ret[y] = append(ret[y], uint8((x+y)/2))
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
