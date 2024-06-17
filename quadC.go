package piscine

import "fmt"

func QuadC(x, y int) {
	if x < 1 || y < 1 {
		return
	}

	x -= 1
	y -= 1
	for i := 0; i <= y; i++ {
		var line string
		for j := 0; j <= x; j++ {
			if (j == 0 && i == 0) || (j == x && i == 0) {
				line += "A"
			} else if ((j == x) && (i == y)) || ((j == 0) && (i == y)) {
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

// j = x , which is left to right
// i = y , which is basically top to bottom
