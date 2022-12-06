package	util

// import "fmt"

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

func (s *Stack) PopN(n int) ([]byte, bool) {
	if len(*s) < n {
		return []byte{}, false
	}
	popped := (*s)[len(*s)-n:]
	// fmt.Println(string(*s))
	if n == len(*s) {
		*s = []byte{}
	} else {
		*s = (*s)[:len(*s)-n]
	}

	// fmt.Println(string(*s))

	return popped, true
}