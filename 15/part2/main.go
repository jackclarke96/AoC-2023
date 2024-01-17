package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("../files/input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	fmt.Println(executeMain(input))
}

func executeMain(bs []byte) int {

	// Split larger input byteSlice to get each individual input as its own byteslice
	bss := splitInputIntoComponents(bs, ',')

	// Maps the lens label to the correct boxNumber.
	hm := HASHMAP{}

	// Map representing the series of 256 boxes
	bm := make(boxMap, 256)

	// Perform algorithm to place each lens in the correct box
	bm.arrangeLensesIntoBoxes(&hm, bss)

	// Now that boxes are correctly placed, calculate total focusing power as described in problem.
	return bm.calculateFocusingPower()

}

func (bm *boxMap) arrangeLensesIntoBoxes(hm *HASHMAP, bss [][]byte) {
	for _, byteSlice := range bss {
		label, boxNumber, action, focalLength := extractLensOperation(hm, byteSlice)
		if action == "=" {
			bm.insertOrUpdateLens(label, boxNumber, focalLength)
		} else {
			bm.removeLensFromBox(label, boxNumber)
		}
	}
}

func (bm *boxMap) calculateFocusingPower() int {
	total := 0
	for boxNumber, box := range *bm {
		for _, lens := range box {
			total += (1 + boxNumber) * lens.focalLength * lens.position
		}
	}
	return total
}

func (bm *boxMap) insertOrUpdateLens(lensLabel label, boxNumber, focalLength int) {
	if existingBox, ok := (*bm)[boxNumber]; ok {
		if existingLens, ok := existingBox[lensLabel]; ok {

			// Both the box and the lens already exist, update the lens to have new focal length
			existingLens.focalLength = focalLength
		} else {

			// The box exists but the lens does not. Create the label at the back of the box
			backOfBox := len(existingBox) + 1
			newLens := lensStruct{focalLength, backOfBox}
			existingBox[lensLabel] = &newLens
		}
	} else {

		// Neither box nor lens exist. Create both the lens and the box
		newLens := lensStruct{focalLength, 1}
		newBox := lensMap{lensLabel: &newLens}
		(*bm)[boxNumber] = newBox
	}
}

func (bm *boxMap) removeLensFromBox(lensLabel label, boxNumber int) {

	// Both box and lens should exist in order for lens to be removed
	if existingBox, ok := (*bm)[boxNumber]; ok {
		if existingLens, ok := existingBox[lensLabel]; ok {

			// Remove the lens and move any other previously behind it forwards
			positionRemoved := existingLens.position
			for _, lens := range existingBox {
				if lens.position > positionRemoved {
					lens.position -= 1
				}
			}
			delete(existingBox, lensLabel)
		}
	}
}

func getBoxNumber(hm *HASHMAP, bs []byte) int {

	// check map to see if already calculated the boxNumber
	if total, ok := (*hm)[string(bs)]; ok == true {
		return total
	}

	// Otherwise, use algorithm described in problem to calculate boxNumber
	total := 0
	for _, b := range bs {
		total += int(b)
		total *= 17
		total = total % 256
	}
	// Append boxNumber to hm to avoid unnecessary recalculation later
	(*hm)[string(bs)] = total

	return total
}
