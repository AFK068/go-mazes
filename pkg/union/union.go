package union

type FindSet struct {
	Parent int
	Rank   int
}

func NewUnionFindSet(x int) *FindSet {
	return &FindSet{
		Parent: x,
		Rank:   0,
	}
}

func Find(sets []*FindSet, x int) int {
	if sets[x].Parent != x {
		sets[x].Parent = Find(sets, sets[x].Parent)
	}

	return sets[x].Parent
}

func Join(sets []*FindSet, x, y int) {
	rootX := Find(sets, x)
	rootY := Find(sets, y)

	if rootX != rootY {
		switch {
		case sets[rootX].Rank > sets[rootY].Rank:
			sets[rootY].Parent = rootX
		case sets[rootX].Rank < sets[rootY].Rank:
			sets[rootX].Parent = rootY
		default:
			sets[rootY].Parent = rootX
			sets[rootX].Rank++
		}
	}
}
