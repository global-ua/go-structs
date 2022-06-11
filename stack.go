package structs

import "errors"

// Stack Taken largely from an example in "Programming In Go"
// Keeping it separate from my stuff
type Stack []interface{}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s *Stack) Top() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("empty stack")
	}
	return (*s)[s.Len()-1], nil
}

func (s *Stack) Pop() (interface{}, error) {
	theStack := *s
	if len(theStack) == 0 {
		return nil, errors.New("empty stack")
	}
	x := theStack[len(theStack)-1]
	*s = theStack[:len(theStack)-1]
	return x, nil
}
