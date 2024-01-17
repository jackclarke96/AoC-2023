package main

// All inputs are contained within one large byteslice, with each input separated by a comma.
// Split this byteslice into each individual input for processing
func splitInputIntoComponents(bs []byte, divider byte) [][]byte {
	startIndex := 0
	bss := [][]byte{}
	for i, b := range bs {
		if b == divider {
			bss = append(bss, bs[startIndex:i])
			startIndex = i + 1
		}
	}
	// account for final string
	bss = append(bss, bs[startIndex:])
	return bss
}

func extractLensOperation(hm *HASHMAP, bs []byte) (label, int, string, int) {
	var l label
	var boxNumber int
	var focalLength int
	var action string

	for i, b := range bs {
		// single quotes compares ascii values
		if b == '-' || b == '=' {
			l = label(bs[:i])
			boxNumber = getBoxNumber(hm, bs[:i])
			action = string(b)
			break
		}
	}

	if action == "=" {
		focalLength = asciiToDigit(bs[len(bs)-1])
	}
	return l, boxNumber, action, focalLength
}

func asciiToDigit(ch byte) int {
	return int(ch) - '0'
}
