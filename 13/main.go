package main

func main() {

}

func executeMain(s string) int {
	return 0
}

func (g grid) getGridScore(horizontalWeight, verticalWeight int) int {
	return verticalWeight*g.getVerticalReflectionLength() + horizontalWeight*g.getHorizontalReflectionLength()
}

func (g grid) getHorizontalReflectionLength() int {
	return 0
}

func (g grid) getVerticalReflectionLength() int {
	return 0
}
