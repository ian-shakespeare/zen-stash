package utils

type Set[T comparable] map[T]bool

func NewSet[T comparable](arr ...T) *Set[T] {
	s := Set[T]{}

	for i := 0; i < len(arr); i += 1 {
		s.Insert(arr[i])
	}

	return &s
}

func (s *Set[T]) Insert(elem T) {
	if v, exists := (*s)[elem]; !exists || !v {
		(*s)[elem] = true
	}
}

func (s *Set[T]) Delete(elem T) {
	(*s)[elem] = false
}

func (s *Set[T]) Array() []T {
	arr := []T{}

	for k, v := range *s {
		if v {
			arr = append(arr, k)
		}
	}

	return arr
}
