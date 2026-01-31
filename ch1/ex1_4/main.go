package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string][]string) // tracks which files contain each line
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, fileNames[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string][]string, name string) {
	input := bufio.NewScanner(f)
	seen := make(map[string]bool)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !seen[line] {
			fileNames[line] = append(fileNames[line], name)
			seen[line] = true
		}
	}
}
