package sets

// String is a set of strings, implemented via map[string]struct{} for minimal memory consumption.
//
// Deprecated: use generic Set instead.
// new ways:
// s1 := Set[string]{}
// s2 := New[string]()
type String map[string]Empty
