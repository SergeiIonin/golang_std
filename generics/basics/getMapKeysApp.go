package main

func main() {

}

func getKeys[K comparable, V any](m map[K]V) []K { // we should specify constraint for generic type
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type customConstraint interface {
	~int | ~string // it means that type itself or the underlying type should be int ot string
}

func getKeys2[K customConstraint, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type Node[T any] struct {
	Val  T
	next *Node[T]
}

func (n *Node[T]) Add(next *Node[T]) {
	n.next = next
}

type SliceFn[T any] struct {
	S       []T
	Compare func(T, T) bool
}

func (s SliceFn[T]) Len() int           { return len(s.S) }
func (s SliceFn[T]) Less(i, j int) bool { return s.Compare(s.S[i], s.S[j]) }
func (s SliceFn[T]) Swap(i, j int)      { s.S[i], s.S[j] = s.S[j], s.S[i] }
