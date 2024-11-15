package domain

type Edge struct {
	First  int
	Second int
}

func NewEdge(f, s int) *Edge {
	return &Edge{f, s}
}
