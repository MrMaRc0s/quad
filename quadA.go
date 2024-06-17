package piscine

import "fmt"

func QuadA(x, y int) {
	if x < 1 || y < 1 {
		return
	}

	x -= 1
	y -= 1
	for i := 0; i <= y; i++ {
		var line string
		for j := 0; j <= x; j++ {
			if (j == 0 && i == 0) || (j == x && i == 0) || (j == 0 && i == y) || (j == x && i == y) {
				line += "o"
			} else if i == 0 || i == y {
				line += "-"
			} else if j == 0 || j == x {
				line += "|"
			} else {
				line += " "
			}
		}
		fmt.Println(line)
	}
}
