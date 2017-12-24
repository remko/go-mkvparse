package mkvparse

import (
	"fmt"
)

func ExampleNameForElementID() {
	fmt.Println(NameForElementID(InfoElement))
	// Output:
	// Info
}
