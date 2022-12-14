package util

type Queue[T any] []T

func (q Queue[T]) IsEmpty() bool {
	return len(q) == 0
}

func (q *Queue[T]) Add(el ...T) {
	*q = append(*q, el...)
}

func (q *Queue[T]) Pop() T {
	var popped T
	if q.IsEmpty() {
		return popped
	}
	popped = (*q)[0]
	*q = (*q)[1:]
	return popped
}
