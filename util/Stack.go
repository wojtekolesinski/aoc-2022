package	util

type Stack []byte

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(el byte) {
	*s = append(*s, el)
}

func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s)-1
	el := (*s)[index]
	*s = (*s)[:index]
	return el, true
}