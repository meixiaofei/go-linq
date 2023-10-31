package linq

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](es ...T) Set[T] {
	s := Set[T]{}
	for _, e := range es {
		s.Add(e)
	}
	return s
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s Set[T]) Add(es ...T) bool {
	flag := true
	for _, e := range es {
		if !s.Contains(e) {
			s[e] = struct{}{}
		} else {
			flag = false
		}
	}
	return flag
}

func (s Set[T]) Remove(es ...T) {
	for _, e := range es {
		delete(s, e)
	}
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s *Set[T]) Clone() Set[T] {
	r := Set[T]{}
	r.Add(s.ToSlice()...)
	return r
}

func (s Set[T]) ToSlice() []T {
	r := make([]T, 0, s.Len())
	for e := range s {
		r = append(r, e)
	}
	return r
}
