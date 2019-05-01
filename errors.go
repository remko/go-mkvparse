package mkvparse

import "fmt"

type ParseError struct {
	Offset int64
	Stack  *ElementStack

	Err error
}

func NewParseError(offset int64, stack *ElementStack, e error) *ParseError {
	return &ParseError{
		Offset: offset,
		Stack:  stack,
		Err:    e,
	}
}

func (pe *ParseError) Error() string {
	return fmt.Sprintf("ParseError: At %d with %s: %s", pe.Offset, pe.Stack, pe.Err)
}

type TruncatedInputError struct {
	Expected int64
	Received int64
}

func (tie *TruncatedInputError) Error() string {
	return fmt.Sprintf("TruncatedInput: Received %d of %d expected bytes.", tie.Received, tie.Expected)
}

type InvalidElementError struct {
	Expected map[ElementID]bool
	Received ElementID
}

func (iee *InvalidElementError) Error() string {
	return fmt.Sprintf("InvalidElement: Received %s but one of %v was expected.", iee.Received, iee.Expected)
}

type InvalidFloatSizeError struct {
	Expected []int64
	Received int64
}

func (ifse *InvalidFloatSizeError) Error() string {
	return fmt.Sprintf("InvalidFloatSize: Received a float size of %d but one of %v was expected.", ifse.Received, ifse.Expected)
}
