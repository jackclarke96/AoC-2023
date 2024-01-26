package main

func convertPipeStart(i, j int, grid *[][]directionChanger) directionChanger {
	startType := replaceStartPipe(i, j, grid)
	return generatePipeStruct(startType)
}

/*
 * Can work out S Pipe type by the values surrounding it.

 * We know that:
 	* (1) each pipe in the loop has exactly 2 connecting pipes.

	* (2) (a) For a pipe to connect to S from the North, it must have a South Direction change
		* (b) For a pipe to connect to S from the East, it must have a West Direction change
		* (c) For a pipe to connect to S from the South, it must have a North Direction change
		* (d) For a pipe to connect to S from the West, it must have an East Direction change

 * Assuming that only one loop is possible (e.g. can't have -S- as well as | above and beneath S) then we can assume only 2 statements from (2) can be true),
   we can check each surrounding entry one by one and eliminate impossible values for S if the adjacent pipe is a possible connection
   e.g. if we have grid[i][j-1] = F, then S cannot be |, L or F, and so on.

   Repeating for each of North, East, South and West will leave only one possible value for S
*/

func replaceStartPipe(i, j int, grid *[][]directionChanger) pipeType {

	var pt pipeType
	var northType, eastType, southType, westType pipeType

	// Get type of pipe directly north
	if i > 0 && (*grid)[i-1][j] != nil {
		northType = (*grid)[i-1][j].getPipe().pipeType
	}

	// Get type of pipe directly east
	if j < len((*grid)[0])-1 && (*grid)[i][j+1] != nil {
		eastType = (*grid)[i][j+1].getPipe().pipeType
	}

	// Get type of pipe south south
	if i < len(*grid)-1 && (*grid)[i+1][j] != nil {
		southType = (*grid)[i+1][j].getPipe().pipeType
	}

	// Get type of pipe directly west
	if j > 0 && (*grid)[i][j-1] != nil {
		westType = (*grid)[i][j-1].getPipe().pipeType
	}

	possibleTypes := map[pipeType]bool{
		NS: true,
		NW: true,
		NE: true,
		SW: true,
		SE: true,
		EW: true,
	}

	// Assuming there will be exactly 2 pipes leading into S, rule out possible types of S
	// based on which S types could connect to the pipe leading into it.
	if northType == SW || northType == SE || northType == NS {
		possibleTypes[SW] = false
		possibleTypes[SE] = false
		possibleTypes[EW] = false
	}
	if eastType == SW || eastType == NW || eastType == EW {
		possibleTypes[NS] = false
		possibleTypes[NW] = false
		possibleTypes[SW] = false
	}
	if southType == NW || southType == NE || southType == NS {
		possibleTypes[NW] = false
		possibleTypes[NE] = false
		possibleTypes[EW] = false
	}
	if westType == SE || westType == NE || westType == EW {
		possibleTypes[NS] = false
		possibleTypes[NE] = false
		possibleTypes[SE] = false
	}

	for key, value := range possibleTypes {
		if value == true {
			pt = key
		}
	}
	return pt
}
