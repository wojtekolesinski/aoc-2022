package	util

// import "fmt"

type Stack[T any] []T

func (s Stack[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack[T]) Push(el T) {
	*s = append(*s, el)
}

func (s *Stack[T]) Pop() (T, bool) {
	var popped T
	if s.IsEmpty() {
		return popped, false
	}
	index := len(*s)-1
	popped = (*s)[index]
	*s = (*s)[:index]
	return popped, true
}

func (s *Stack[T]) PopN(n int) ([]T, bool) {
	if len(*s) < n {
		return []T{}, false
	}
	popped := (*s)[len(*s)-n:]
	if n == len(*s) {
		*s = []T{}
	} else {
		*s = (*s)[:len(*s)-n]
	}


	return popped, true
}