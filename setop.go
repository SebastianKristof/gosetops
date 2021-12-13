package setop

type Set map[string]struct{}

// NewSet
func NewSet(elements ...string) Set {
	s := make(Set)
	s.Add(elements...)

	return s
}

// Includes
func (s Set) Includes(str string) bool {
	_, ok := s[str]

	return ok
}

// IncludesSet
func (s Set) IncludesSet(s2 Set) bool {
	for k := range s2 {
		if _, ok := s[k]; !ok {
			return false
		}
	}

	return true
}

// EqualTo
func (s Set) EqualTo(s2 Set) bool {
	if len(s) != len(s2) {
		return false
	}

	if !s.IncludesSet(s2) {
		return false
	}

	return true
}

// LargerThan
func (s Set) LargerThan(s2 Set) bool {
	if len(s) <= len(s2) {
		return false
	}

	if !s.IncludesSet(s2) {
		return false
	}

	return true
}

// SmallerThan
func (s Set) SmallerThan(s2 Set) bool {
	if len(s) >= len(s2) {
		return false
	}

	if !s2.IncludesSet(s) {
		return false
	}

	return true
}

// Add
func (s Set) Add(elements ...string) {
	for _, e := range elements {
		s[e] = struct{}{}
	}
}

// Remove

// RemoveMany ?

// Union

// Intersection

// Difference
