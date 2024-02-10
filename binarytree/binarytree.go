package binarytree

type Tree interface {
	Search(value int) *knot
	Insert(value int)
	String() string
	GiveKnots() []*knot
	InsertList(values []int)
}
