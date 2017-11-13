// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	//unique word maps to map that maps filenames to counts 
	counts := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		total := 0
		for _, i := range n {
			total += i
		}

		if total > 1 {
			fmt.Printf("%s :\t\t%d\t",line,total)
			for fn,_ := range n{
				fmt.Printf("%s ",fn)
			}
			fmt.Printf("\n")
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int, fn string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		q := counts[input.Text()]

		if q == nil{
			q = make(map[string]int)
			counts[input.Text()] = q
		}
		q[fn]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
