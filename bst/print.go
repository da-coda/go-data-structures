package bst

import (
	"fmt"
	"strconv"
	"strings"
)

func (N node) Print(space int, spaceSize int) {
	if N.HasRight() {
		N.right.Print(space+spaceSize, N.Height())
	}
	fmt.Println(strings.Repeat(" ", space) + strconv.Itoa(N.key))
	if N.HasLeft() {
		N.left.Print(space+spaceSize, N.Height())
	}
}
