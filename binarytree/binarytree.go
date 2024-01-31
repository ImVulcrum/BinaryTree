package binarytree

type Tree interface {
	Search(value int) *knot
	Insert(value int)
	Print(terminal_width int) string
	GiveKnots() []*knot
	InsertList(values []int)
}
