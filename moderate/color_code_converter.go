package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func abs(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}

func mod2(a float32) float32 {
	return a - float32(int(a/2)*2)
}

func rint(a float32) int {
	return int(a*255 + 0.5)
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var r, g, b int
		line := scanner.Text()
		switch {
		case line[0] == '#':
			fmt.Sscanf(line, "#%2x%2x%2x", &r, &g, &b)
		case line[:3] == "HSL":
			var h, s, l float32
			fmt.Sscanf(line, "HSL(%f,%f,%f)", &h, &s, &l)
			s, l = s/100, l/100
			c, h2 := (1-abs(2*l-1))*s, h/60
			x, m := c*(1-abs(float32(mod2(h2)-1))), l-c/2
			switch {
			case h2 < 1:
				r, g, b = rint(c+m), rint(x+m), rint(m)
			case h2 < 2:
				r, g, b = rint(x+m), rint(c+m), rint(m)
			case h2 < 3:
				r, g, b = rint(m), rint(c+m), rint(x+m)
			case h2 < 4:
				r, g, b = rint(m), rint(x+m), rint(c+m)
			case h2 < 5:
				r, g, b = rint(x+m), rint(m), rint(c+m)
			case h2 < 6:
				r, g, b = rint(c+m), rint(m), rint(x+m)
			}
		case line[:3] == "HSV":
			var h, s, v float32
			fmt.Sscanf(line, "HSV(%f,%f,%f)", &h, &s, &v)
			s, v = s/100, v/100
			c, h2 := v*s, h/60
			x, m := c*(1-abs(float32(mod2(h2)-1))), v-c
			switch {
			case h2 < 1:
				r, g, b = rint(c+m), rint(x+m), rint(m)
			case h2 < 2:
				r, g, b = rint(x+m), rint(c+m), rint(m)
			case h2 < 3:
				r, g, b = rint(m), rint(c+m), rint(x+m)
			case h2 < 4:
				r, g, b = rint(m), rint(x+m), rint(c+m)
			case h2 < 5:
				r, g, b = rint(x+m), rint(m), rint(c+m)
			case h2 < 6:
				r, g, b = rint(c+m), rint(m), rint(x+m)
			}
		case line[0] == '(':
			var c, m, y, k float32
			fmt.Sscanf(line, "(%f,%f,%f,%f)", &c, &m, &y, &k)
			r, g, b = rint((1-c)*(1-k)), rint((1-m)*(1-k)), rint((1-y)*(1-k))
		}
		fmt.Printf("RGB(%d,%d,%d)\n", r, g, b)
	}
}
