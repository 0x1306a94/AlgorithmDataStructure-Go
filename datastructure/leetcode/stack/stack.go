package stack

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack) Cap() int {
	return cap(s)
}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s Stack) Top() interface{} {
	if len(s) == 0 {
		return nil
	}
	return s[len(s)-1]
}

func (s *Stack) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	thenStack := *s
	oldValue := thenStack[len(thenStack)-1]
	*s = thenStack[:len(thenStack)-1]
	return oldValue
}
