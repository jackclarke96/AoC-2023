package main

func NewPipeStart(i, j int, matrix [][]DirectionChanger) DirectionChanger {
	startType := replaceStartPipe(i, j, matrix)
	switch startType {
	case NS:
		return NewPipeNS(i, j)
	case EW:
		return NewPipeEW(i, j)
	case NE:
		return NewPipeNE(i, j)
	case NW:
		return NewPipeNW(i, j)
	case SW:
		return NewPipeSW(i, j)
	case SE:
		return NewPipeSE(i, j)
	}
	return nil
}

func replaceStartPipe(i, j int, matrix [][]DirectionChanger) PipeType {

	var pt PipeType
	var northType, eastType, southType, westType PipeType

	if i > 0 && matrix[i-1][j] != nil {
		northType = matrix[i-1][j].GetPipe().Type
	}
	if j < len(matrix[0])-1 && matrix[i][j+1] != nil {
		eastType = matrix[i][j+1].GetPipe().Type
	}
	if i < len(matrix)-1 && matrix[i+1][j] != nil {
		southType = matrix[i+1][j].GetPipe().Type
	}
	if j > 0 && matrix[i][j-1] != nil {
		westType = matrix[i][j-1].GetPipe().Type
	}

	possibleTypes := map[PipeType]bool{
		NS: true,
		NW: true,
		NE: true,
		SW: true,
		SE: true,
		EW: true,
	}

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
