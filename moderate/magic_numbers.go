package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isMagic(a int) bool {
	var dig, r uint
	ns := []uint{}
	for a > 0 {
		r = uint(a % 10)
		if r == 0 || dig&(1<<r) > 0 {
			return false
		}
		dig |= 1 << r
		ns = append(ns, r)
		a /= 10
	}
	dig, r = 0, 0
	for i := 0; i < len(ns); i++ {
		r = (r + ns[(uint(len(ns))-1-r)]) % uint(len(ns))
		if dig&(1<<r) > 0 {
			return false
		}
		dig |= 1 << r
	}
	return r == 0
}

var magic []int

func init() {
	magic = []int{}
	for i := 1; i <= 9876; i++ {
		if isMagic(i) {
			magic = append(magic, i)
		}
	}
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var a, b int
		fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
		r := []string{}
		for i := 0; i < len(magic) && magic[i] <= b; i++ {
			if magic[i] >= a {
				r = append(r, fmt.Sprint(magic[i]))
			}
		}
		if len(r) > 0 {
			fmt.Println(strings.Join(r, " "))
		} else {
			fmt.Println(-1)
		}
	}
}
