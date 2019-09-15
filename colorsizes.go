package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	numberRE = regexp.MustCompile(`^\d+`)
	palette  = []string{"\033[0;1;30m", "\033[0m", "\033[0;1m", "\033[0;1;31m", "\033[0;1;33m"}
)

func splitIntoThrees(s string) []string {
	n := (len(s) + 2) / 3
	parts := make([]string, 0, n)
	for i := (len(s)+2)%3 + 1; s != ""; s, i = s[i:], 3 {
		parts = append(parts, s[:i])
	}
	return parts
}

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			return
		}
		if i := strings.IndexFunc(line, func(r rune) bool {
			return !('0' <= r && r <= '9')
		}); i != 0 {
			if i == -1 {
				i = len(line)
			}
			colored := ""
			colors := palette
			parts := splitIntoThrees(line[:i])
			for i := len(parts) - 1; i >= 0; i-- {
				colored = colors[0] + parts[i] + colored
				if len(colors) > 1 {
					colors = colors[1:]
				}
			}
			fmt.Printf("%s\033[0m%s", colored, line[i:])
		} else {
			fmt.Println(line)
		}
	}
}
