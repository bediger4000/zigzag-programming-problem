# Daily Coding Problem: Problem #521 [Medium]  

Let's beat another Daily Coding Problem to death!

### Problem Statement

Given a string and a number of lines k, print the string in zigzag form. In
zigzag, characters are printed out diagonally from top left to bottom right
until reaching the kth line, then back up to top right, and so on.

For example, given the sentence "thisisazigzag" and k = 4, you should print:

    t     a     g
     h   s z   a
      i i   i z
       s     g

I'm doing this in Go, which has the advantage of the `range` built-in that
enourages less mistake-prone array index and value retrieval.

Go uses Unicde strings in UTF-8 format, so each input string gets normalized
so that seperate composable code points get composed.

## Build and run

    $ go build method1.go
    $ ./method1 4 thisisazigzag
    t     a     g
     h   s z   a
      i i   i z
       s     g

    $ go build method2.go
    $ ./method2 4 thisisazigzag
    t     a     g
     h   s z   a
      i i   i z
       s     g

## Method 1

Let `l` hold the length of the input string, `k` is the height of the zigzag.

Create an array of array-of-runes `k` long.
Create each array-of-runes to be `l` long,
and put blank characters at each index.

Walk the input string's runes, each at index `i`.
Start with a height of 0, and a direction up.
Put the rune `i` in the array of array-of-runes at `block[i][height]`.
If the height has zero value, or has the value `k - 1`, reverse
direction.

I cleverly let "up" direction be an int 1 and "down" be an int -1,
so I could merely `+=` the height with the direction.
The value of `height` varies between 0 and `k - 1`.
It's essentially a function of the index in the input string.

After walking the string, print out each of the arrays-of-runes,
from 0 to `k - 1`.
Even though we started at height of 0,
the 0-index array-of-runes is the top of the zigzag as printed out.

## Method 2

Make `k` passes over the runes of the input string,
keeping track of the "height" in the zigzag each input rune should lie at.
This is identical to keeping track of the height in method 1.
If the height of the rune equals the pass number,
print out the rune.
Otherwise print out a space character.

## Analysis

|Factor       |Advantage    |
|-------------|-------------|
|Memory       | method 2    |
|Input string access| method 1|
|Trailing blanks| tie |
|Cognitive Complexity| method 1|

Method 1 allocates a potentially large block of heap.
Method 2 just prints things out on the fly.

Method 1 makes a single pass over the input string,
method 2 makes `k` passes.

It looks to me like you trade the space used by method 1
for the multiple passes used by method 2.
A classic time vs space tradeoff.

As I implemented them, both methods print out trailing blanks,
blanks on upper lines that are "above" the height of the last upward zag,
or below the last downward zig.
Both of them could implement something to prevent this,
like the Go standard library function `strings.TrimRight` on method 1, a simple check on row index for method 2.

The "cognitive complexity" I cite is pretty informal.
It seems to me it's harder to get method 2 correct.
I wrote the method 1 code first, and only after seeing how to implement "up and down"
did I consider method 2 a possiblity.
