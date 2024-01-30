package binarytree


type Treee interface {
	Search (value int) *Knot
	Insert  (value int)
	String () string
	SubSearch(value int) *Knot
	GiveKnots() []*Knot
}
