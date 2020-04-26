package main

/*
Given a string and a number of lines k, print the string in zigzag form. In
zigzag, characters are printed out diagonally from top left to bottom right
until reaching the kth line, then back up to top right, and so on.

For example, given the sentence "thisisazigzag" and k = 4, you should print:

t     a     g
 h   s z   a
  i i   i z
   s     g

*/

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/text/unicode/norm"
)

/*
 * Method 2: traverse the input string one rune at a time,
 * n-1 (height of zig-zag) times.
 * For each traverse of the input string, calculate a "height"
 * for each rune. If that height numerically equates to the traverse
 * we're in, print out the rune. Otherwise, print out a blank.
 * Print a newline at the end of each traverse.
 *
 */

const (
	down int = 1
	up   int = -1
)

func main() {
	// k is height of zig-zag.
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Compose any random accents that are "after" the letter
	// they compose with. This probably still isn't good enough
	// if you're doing bidirectional text.
	runes := []rune(norm.NFC.String(os.Args[2]))

	for j := 0; j < k; j++ {
		height := 0
		dir := up
		for _, r := range runes {
			out := ' '
			if height == j {
				out = r
			}
			fmt.Printf("%c", out)
			if height == 0 || height == (k-1) {
				dir = -dir
			}
			height += dir
		}
		fmt.Println()
	}
}
