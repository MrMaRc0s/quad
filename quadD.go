package piscine

import "fmt"

func QuadD(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	x -= 1
	y -= 1
	for i := 0; i <= y; i++ {
		var line string
		for j := 0; j <= x; j++ {
			if (j == 0 && i == 0) || (j == 0 && i == y) {
				line += "A"
			} else if ((j == x) && (i == y)) || ((j == x) && (i == 0)) {
				line += "C"
			} else if i == 0 || i == y {
				line += "B"
			} else if j == 0 || j == x {
				line += "B"
			} else {
				line += " "
			}
		}
		fmt.Println(line)
	}
}
