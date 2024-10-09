package domain

type Edge struct {
	first  int
	second int
}

func NewEdge(f, s int) *Edge {
	return &Edge{f, s}
}

func (e *Edge) GetFirst() int {
	return e.first
}

func (e *Edge) GetSecond() int {
	return e.second
}
