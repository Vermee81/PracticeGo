package main

import "tour/pic"

func Pic(dx, dy int) [][]uint8 {
	graph := make([][]uint8, dy)
	for x := range graph {
		graph[x] = make([]uint8, dx)
		for y := range graph[x] {
			graph[x][y] = uint8(x^y)
		}
	}
	return graph
}

func main() {
	pic.Show(Pic)
}