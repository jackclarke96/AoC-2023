## Attempt 1

Created this function:

```go
package main

import (
	"fmt"
	"regexp"
	"strings"
)

var numberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// returns a string of numbers e.g. 1234
func stringToNumbers(s string) string {
	newStr := strings.ToLower(s)
	fmt.Println("lower case string", newStr)
	for word, number := range numberMap {
		re := regexp.MustCompile(`\b` + word + `\b`)
		newStr = re.ReplaceAllString(newStr, number)
		// newStr = strings.ReplaceAll(newStr, word, number)
	}
	fmt.Println("new string", newStr)

	return newStr
}

```

However, this leads to inconsistencies due the following:

* The word 'twone' could either be 2 or 1, depending on which key is evaluated first. Therefore need to iterate through the strings rather than the map.
* maps have a randomised iteration order


To guarantee a specific iteration order, create some additional data structures and use a method from the sort package to control the order of iteration e.g.

```go
func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// Create a struct for each map key-value pair
	// Make sure the types match accordingly
	type KeyValue struct {
		Key   string
		Value int
	}

	// create an empty slice of key-value pairs
	s := make([]KeyValue, 0, len(m))
	// append all map keys-value pairs to the slice
	for k, v := range m {
		s = append(s, KeyValue{k, v})
	}
}
```

### Alternative method

Can no longer use a map to deal with this. Instead create two substrings using for loops. 

Work through each string character by character from left to right and add eachc character to a substring. Evaluate the substring on every iteration for a number or string number.

The second substring will be evaluated in exactly the same way but we move right to left through the original string.

```go
// Iterate forwards through the string
for i := 0; i < len(s); i++ {
    mySubstr += string(s[i])
    stringMatch := reWords.FindString(mySubstr)
    if stringMatch != "" {
        return numberMap[stringMatch]
    }
    numberMatch := reDigits.FindString(mySubstr)
    if numberMatch != "" {
        return numberMatch
    }
}
```

Then clean up - pull dupe logic out and use some GoIsh syntax e.g. combining if statements with the assignment like this

```go
if num := findFirstNumberInSubstring(mySubstr, reWords, reDigits); num != "" {
            return num
        }
```

Get the final thing by also realising that don't need to iterate through backwards if no match is found in forwards direction.