package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func pos(l, r uint8) string {
	return string('a'+l) + string('1'+r)
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	m := []string{}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		l, r := scanner.Text()[0]-'a', scanner.Text()[1]-'1'
		if l > 1 && r > 0 {
			m = append(m, pos(l-2, r-1))
		}
		if l > 1 && r < 7 {
			m = append(m, pos(l-2, r+1))
		}
		if l > 0 && r > 1 {
			m = append(m, pos(l-1, r-2))
		}
		if l > 0 && r < 6 {
			m = append(m, pos(l-1, r+2))
		}
		if l < 7 && r > 1 {
			m = append(m, pos(l+1, r-2))
		}
		if l < 7 && r < 6 {
			m = append(m, pos(l+1, r+2))
		}
		if l < 6 && r > 0 {
			m = append(m, pos(l+2, r-1))
		}
		if l < 6 && r < 7 {
			m = append(m, pos(l+2, r+1))
		}
		fmt.Println(strings.Join(m, " "))
		m = []string{}
	}
}
