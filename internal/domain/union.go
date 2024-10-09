package domain

type UnionFindSet struct {
	parent int
	rank   int
}

func NewUnionFindSet(x int) *UnionFindSet {
	return &UnionFindSet{
		parent: x,
		rank:   0,
	}
}
