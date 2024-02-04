// package main

// func checkRemainingSpringsInBounds(i, iMax, numberSpringsRemaining, totalSpringsLength int) bool {
// 	minimumLengthWithGaps := totalSpringsLength + numberSpringsRemaining - 1
// 	if i+minimumLengthWithGaps < iMax {
// 		return true
// 	}
// 	return false
// }

// func checkProposedSpringInBounds(j, jMax, i, iMax, springLength int) bool {
// 	if j < jMax {
// 		// check that both the hashes and the dot following it will fit. No -1 needed. First # already accounted for but also have . at end so evens out
// 		return i+springLength <= iMax
// 	}
// 	// check that just the hashes will fit. -1 because we are already at first hash
// 	return i+springLength-1 <= iMax
// }

/*
User
func generateCombinations(springString []string, springLengths []int) int {
	questionMarkIndices := findQuestionMarkIndices(springString)
	iMaxStart := calculateIMax(springString, springLengths)
	memo := map[StateKey]int{}
	total := 0

	numberOfHashes := 5

	var closure func(depth int) int

	closure = func(depth int) int {
		if depth == len(questionMarkIndices) {
			total += isValidCombination(springLengths, springString, iMaxStart)
			fmt.Println(total)
			return total
		}
		for _, elem := range []string{".", "#"} {
			springString[questionMarkIndices[depth]] = elem
			if !isValidPartial(springLengths, springString[:depth+1], iMaxStart) {
				return
			}

			if val, ok := memo[StateKey{depth, 3}]; ok { // replace 3 with number of hashes
				total += val
			} else {
				memo[StateKey{depth, 3}] = closure(depth + 1)
			}

			// if (springsPlacedAt this point seen before) {
			// myMap += valueWhenExpored (sum it on again)
			// myMap

			// myMap.valueWhenExplored = closure(depth+1)
			closure(depth + 1)
		}
	}
	closure(0)
	return total
}

var memo = make(map[StateKey]int)

var closure func(depth int, numberOfHashes int) int
closure = func(depth int, numberOfHashes int) int {
    key := StateKey{Depth: depth, NumHashes: numberOfHashes}

    // Check if this state has been computed before
    if val, found := memo[key]; found {
        return val // Use the memoized value
    }

    // Base case: If all question marks have been replaced
    if depth == len(questionMarkIndices) {
        result := isValidCombination(springLengths, springString, iMaxStart)
        memo[key] = result // Memoize the result before returning
        return result
    }

    localTotal := 0 // Local total for this depth
    for _, elem := range []string{".", "#"} {
        springString[questionMarkIndices[depth]] = elem
        if isValidPartial(springLengths, springString[:depth+1], iMaxStart) {
            // Accumulate results from deeper recursive calls
            localTotal += closure(depth+1, calculateNumberOfHashes(springString[:depth+1]))
        }
    }

    memo[key] = localTotal // Memoize the computed total for this state
    return localTotal
}

total = closure(0, calculateNumberOfHashes(springString))
return total



func generateCombinations(springString []string, springLengths []int) int {
	questionMarkIndices := findQuestionMarkIndices(springString)
	iMaxStart := calculateIMax(springString, springLengths)
	memo := make(map[StateKey]int)
	total := 0

	var closure func(depth int, numHashes int, previousChar string) int
	closure = func(depth int, numHashes int, previousChar string) int {
		// Generate the key based on current state
		key := StateKey{Depth: depth, NumHashes: numHashes, PreviousChar: previousChar}

		// Check if we have already computed this state
		if val, found := memo[key]; found {
			return val // Return the memoized value
		}

		// Base case: if all question marks have been processed
		if depth == len(questionMarkIndices) {
			// fmt.Println(key)
			// fmt.Println(springString)
			result := isValidCombination(springLengths, springString, iMaxStart)
			// memo[key] = result // Store the result in the memoization map
			return result
		}

		localTotal := 0 // Local total for combinations from this state forward
		for _, elem := range []string{".", "#"} {
			springString[questionMarkIndices[depth]] = elem
			updatedNumHashes := numHashes
			if elem == "#" {
				updatedNumHashes++
			}
			if isValidPartial(springLengths, springString[:depth+1], iMaxStart) {
				// Recurse and accumulate results
				localTotal += closure(depth+1, updatedNumHashes, elem)
			}
		}

		memo[key] = localTotal // Memoize the computed total for this state
		return localTotal
	}

	// Initial call to the closure with starting values
	total = closure(0, 0, "") // Assuming "" as the initial previous character
	return total
}


func generateCombinations(springString []string, springLengths []int) int {
	questionMarkIndices := findQuestionMarkIndices(springString)
	iMaxStart := calculateIMax(springString, springLengths)
	memo := map[StateKey]int{}
	total := 0

	numberOfHashes := 5

	var closure func(depth int) int

	closure = func(depth int) int {
		if depth == len(questionMarkIndices) {
			total += isValidCombination(springLengths, springString, iMaxStart)
			fmt.Println(total)
			return total
		}
		for _, elem := range []string{".", "#"} {
			springString[questionMarkIndices[depth]] = elem
			if !isValidPartial(springLengths, springString[:depth+1], iMaxStart) {
				return
			}

			if val, ok := memo[StateKey{depth, 3}]; ok { // replace 3 with number of hashes
				total += val
			} else {
				memo[StateKey{depth, 3}] = closure(depth + 1)
			}

			// if (springsPlacedAt this point seen before) {
			// myMap += valueWhenExpored (sum it on again)
			// myMap

			// myMap.valueWhenExplored = closure(depth+1)
			closure(depth + 1)
		}
	}
	closure(0)
	return total
}
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1 := flag.Bool("p1", false, "run part 1")
	p2 := flag.Bool("p2", false, "run part 2")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("no input file provided")
	}

	b, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	input := string(b)

	if *p1 {
		fmt.Println("part 1:", part1(input))
	}

	if *p2 {
		fmt.Println("part 2:", part2(input))
	}
}

func part1(input string) int {
	records, groups := parse(input)
	res := 0
	for i := range records {
		res += solve(records[i], groups[i])
	}

	return res
}

func part2(input string) int {
	records, groups := parse(input)
	res := 0
	for i := range records {
		res += solve(unfoldRecord(records[i]), unfoldGroup(groups[i]))
	}

	return res
}

func unfoldRecord(record string) string {
	var res strings.Builder
	for i := 0; i < len(record)*5; i++ {
		if i != 0 && i%len(record) == 0 {
			res.WriteByte('?')
		}
		res.WriteByte(record[i%len(record)])
	}

	return res.String()
}

func unfoldGroup(group []int) []int {
	var res []int
	for i := 0; i < len(group)*5; i++ {
		res = append(res, group[i%len(group)])
	}

	return res
}

func solve(record string, group []int) int {
	var cache [][]int
	for i := 0; i < len(record); i++ {
		cache = append(cache, make([]int, len(group)+1))
		for j := 0; j < len(group)+1; j++ {
			cache[i][j] = -1
		}
	}

	return dp(0, 0, record, group, cache)
}

func dp(i, j int, record string, group []int, cache [][]int) int {
	if i >= len(record) {
		if j < len(group) {
			return 0
		}
		return 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	res := 0
	if record[i] == '.' {
		res = dp(i+1, j, record, group, cache)
	} else {
		if record[i] == '?' {
			res += dp(i+1, j, record, group, cache)
		}
		if j < len(group) {
			count := 0
			for k := i; k < len(record); k++ {
				if count > group[j] || record[k] == '.' || count == group[j] && record[k] == '?' {
					break
				}
				count += 1
			}

			if count == group[j] {
				if i+count < len(record) && record[i+count] != '#' {
					res += dp(i+count+1, j+1, record, group, cache)
				} else {
					res += dp(i+count, j+1, record, group, cache)
				}
			}
		}
	}

	cache[i][j] = res
	return res
}

func parse(input string) ([]string, [][]int) {
	var records []string
	var groups [][]int

	for _, line := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n") {
		parts := strings.Split(line, " ")
		records = append(records, parts[0])
		var group []int
		for _, num := range strings.Split(parts[1], ",") {
			num, _ := strconv.Atoi(num)
			group = append(group, num)
		}
		groups = append(groups, group)
	}

	return records, groups
}
