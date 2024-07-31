package models

// https://gist.github.com/bgadrian/cb8b9344d9c66571ef331a14eb7a2e80
type Set struct {
	list map[string]struct{} // empty structs occupy 0 memory
}

func (s *Set) Has(v string) bool {
	_, ok := s.list[v]
	return ok
}

func (s *Set) Add(v string) {
	s.list[v] = struct{}{}
}

func NewSet() *Set {
	return &Set{
		list: make(map[string]struct{}),
	}
}

func (s *Set) ToSlice() []string {
	keys := make([]string, 0, len(s.list))
	for k := range s.list {
		keys = append(keys, k)
	}
	return keys
}
