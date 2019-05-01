package mkvparse

import "strings"

// A stack of elements.
type ElementStack []Element

func NewElementStack() *ElementStack {
	return &ElementStack{}
}

// Push the id and info onto the element stack.
func (es *ElementStack) Push(id ElementID, info ElementInfo) {
	e := Element{
		ID:   id,
		Info: info,
	}

	*es = append(*es, e)
}

// Remove the top element on the stack and return it.
func (es *ElementStack) Pop() (id ElementID, info ElementInfo) {
	l := len(*es)

	e := (*es)[l-1]
	*es = (*es)[:l-1]

	return e.ID, e.Info
}

// Return the top element on the stack.
func (es *ElementStack) Peek() (id ElementID, info ElementInfo) {
	l := len(*es)
	e := (*es)[l-1]

	return e.ID, e.Info
}

// Return the index of the last element with an unknown size. If no element of
// unknown size exists, then index will be 0 and ok will be false.
func (es *ElementStack) UnknownSize() (index int, ok bool) {
	for i, e := range *es {
		if e.Info.Size == UnknownSize {
			index = i
			ok = true
		}
	}

	return index, ok
}

// Render a human readable list of stack id names.
func (es *ElementStack) String() string {
	ids := []string{}

	for _, e := range *es {
		ids = append(ids, NameForElementID(e.ID))
	}

	return "[" + strings.Join(ids, ",") + "]"
}
