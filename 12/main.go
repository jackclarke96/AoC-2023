package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/combin"
)

func main() {

}

func executeMain(s string) int {
	return 5
}

func getCombinations(s string, brokenOrder []int) int {
	fmt.Println(s)
	fmt.Println(generateSubstrSlice(brokenOrder))
	combin.Binomial(5, 2)
	return 5
}

func convertToSubstr(length int, addSpace bool) string {
	if length <= 0 {
		return ""
	}
	str := ""
	for i := 0; i < length; i++ {
		str += "#"
	}
	if addSpace {
		str += "."
	}
	return str
}

type damagedSpringGroup map[int]string

func generateSubstrSlice(brokenOrder []int) damagedSpringGroup {
	mapping := make(damagedSpringGroup, len(brokenOrder))
	for i, val := range brokenOrder {
		if i == len(brokenOrder)-1 {
			mapping[i] = convertToSubstr(val, false)
		} else {
			mapping[i] = convertToSubstr(val, true)
		}
	}
	return mapping
}

// plan. Find string slices that MUST go somewhere then remove that part of the string.
// after that, use combinatorics

/*
????.######..#####. 4,1,1 answer iks 4
[#. ######. #####] ===> we must have *#.*######.*#####* and our string is of length 18. We have a string of length2, a string of length 7, and a string of length 5 = 14 spaces spoken for.


I could create my string #.######.##### of length 14, then insert '.'s into *#.*######.*#####*
But then we know that string[17] is a dot.
We also know that string[11] and string[10] are dots Can i use this?

If i replace every . that follows immediately after with another symbol, does that simplify things? since these absolutely have to be there.

I could create a chain of some sort [wildcard1, "#.", wildcard2, "######.", wildcard3, "#####", wildcard4].
This will always be of length 2x(number of contiguous groups of damaged springs) + 1.
The number of wildcards is number of contiguous groups of damaged springs + 1,
and the number of dots to fit in those wildcards is length strLength - (number of damaged springs + number of contiguous groups of damaged springs -1)

if we have a string of length 18, with four wildcards, filling in the wildcards gives

then fill in the wildcards we know to create a slice. How though???

[wildcard1, "#.", wildcard2, "######.", wildcard3, "#####", "wildcard4""]

fill in the wild cards with

????.######..#####.
[]
*/

/*
?###???????? with 3,2,1

*###.*##.*#*

String has length 12.

8 spaces spoken for by ###., ##., #

4 spaces in which we must distribute 4 "."s, AND  index 1,2,3 all hashes.

Do we just insert the .s?  how many combinations would we need to try

Breaking replacements into cases:
* When replacing a ? before a #:
	* If we replace with a # we get a string of length 1+length 0f #s
	* If we
	Could split the actual string on dots too
*/
