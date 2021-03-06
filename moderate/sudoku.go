package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func sq(a int) (ret int) {
	switch a {
	case 4:
		ret = 2
	case 9:
		ret = 3
	}
	return ret
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ";")
		t := strings.Split(s[1], ",")
		var n, d int
		fmt.Sscan(s[0], &n)
		valid := true
		csqu, col := make([]int, sq(n)), make([]int, n)
		for i := 0; valid && i < n; i++ {
			var crow int
			for j := 0; j < n; j++ {
				fmt.Sscan(t[i*n+j], &d)
				d = 1 << uint(d-1)
				crow |= d
				csqu[j/sq(n)] |= d
				col[j] |= d
			}
			d = (1 << uint(n)) - 1
			if crow != d {
				valid = false
				break
			}
			if ((i + 1) % sq(n)) == 0 {
				for j := 0; j < sq(n); j++ {
					if csqu[j] != d {
						valid = false
						break
					} else {
						csqu[j] = 0
					}
				}
			}
		}
		for i := 0; valid && i < n; i++ {
			if col[i] != (1<<uint(n))-1 {
				valid = false
			}
		}
		if valid {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
