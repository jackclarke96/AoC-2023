
1. Take maximum of each colour across each game and add it to a map
2. multiply the maximums

```go
type colourMap map[string]int

func getMaximumOfEachColour(coloursDrawn []string, numberOfTimes []string) colourMap {

	maxMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for j, colour := range coloursDrawn {
		timesDrawn, err := strconv.Atoi(numberOfTimes[j])
		if err != nil {
			log.Println("Error extracting number")
			continue
		}
		if (maxMap[colour]) < timesDrawn {
			maxMap[colour] = timesDrawn
		}
	}
	return maxMap
}

func (c colourMap) getCubes() int {
	fmt.Println(c)
	fmt.Println(c["red"] * c["green"] * c["blue"])
	return c["red"] * c["green"] * c["blue"]
}
```

