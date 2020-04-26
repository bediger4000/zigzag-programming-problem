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
 * Method 1: allocate n []runes, each of same length as input string.
 * Fill each n []runes with blank characters.
 * Set up to walk the string: set direction to "down". Set vertical
 * index to 0.
 * For each rune at position i in the input string,
 * choose rune array based on vertical index.
 * put the rune in index i in that []rune.
 * Increment i.
 * Check if vertical index is at a limit, 0 or n - 1.
 * If so, change direction. Direction numerically represented
 * as 1 (up) or -1 (down) so that increment of j (row number) changes
 * j correctly, zigging down or zagging up.
 * When you reach the end of input string, print out the rows, 0 through
 * n - 1 in order, so it looks like top-to-bottom.
 */

const (
	down int = 1
	up   int = -1
)

func main() {
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	runes := []rune(norm.NFC.String(os.Args[2]))
	l := len(runes)

	block := make([][]rune, k)
	for i := range block {
		block[i] = make([]rune, l)
		for j := range block[i] {
			block[i][j] = ' '
		}
	}

	direction := up
	for i, j := 0, 0; i < l; i++ {
		block[j][i] = runes[i]
		if j == (k-1) || j == 0 {
			direction = -direction
		}
		j += direction
	}

	for _, row := range block {
		fmt.Printf("%s\n", string(row))
	}

	fmt.Printf("\n")
}
