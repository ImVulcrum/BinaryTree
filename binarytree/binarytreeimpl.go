package binarytree

import "fmt"

type tree struct{
	knots []*Knot
	root *Knot
	current_knot *Knot
	knot_count int
}

type Knot struct{
	leftPtr *Knot
	rightPtr *Knot
	value int
}


func NewKnot(value int) *Knot{
	var k *Knot = new(Knot)
	k.value = value
	k.leftPtr = nil
	k.rightPtr = nil
	return k
}

func NewTree() *tree{
	var t *tree = new(tree)
	return t
}

func (t *tree) GiveKnots () []*Knot {
	return t.knots
}


func (t *tree) Insert (value int) {
	var knot *Knot = NewKnot(value)
	t.knot_count++
	t.knots = append(t.knots, knot)

	if t.root == nil {
		t.root = knot
		return
	}
	t.current_knot = t.root
	for {
		if value <= t.current_knot.value && t.current_knot.leftPtr != nil{
			t.current_knot = t.current_knot.leftPtr
		} else if t.current_knot.rightPtr != nil {
			t.current_knot = t.current_knot.rightPtr
		}	else {
			break
		}
	}

	if value <= t.current_knot.value {
		t.current_knot.leftPtr = knot
	}	else if value > t.current_knot.value {
		t.current_knot.rightPtr = knot
	}
}

func (n *tree) SubSearch (value int) *Knot {
	if n.current_knot == nil {
		return nil
	} else if value < n.current_knot.value {
		n.current_knot = n.current_knot.leftPtr
		return n.SubSearch(value)
	} else if value > n.current_knot.value {
		n.current_knot = n.current_knot.rightPtr
		return n.SubSearch(value)
	} else {
		return n.current_knot
	}
}


func (n* tree) Search(value int) *Knot {
	n.current_knot = n.root
	knot := n.SubSearch(value)
	return knot 
}

func (current_knot *Knot) AddTwoKnots(s string) (string, *Knot, *Knot) {
	var left_knot string = "nil"
	var right_knot string = "nil"

	if current_knot.leftPtr != nil {
		left_knot = fmt.Sprint(current_knot.leftPtr.value)
	}
	if current_knot.rightPtr != nil {
		right_knot = fmt.Sprint(current_knot.rightPtr.value)
	}
	s = s + " " + left_knot +  " " + right_knot + " "
	return s, current_knot.leftPtr, current_knot.rightPtr
	}

func interate_trough_layer (s string, layer []*Knot) (string, []*Knot) {
	var new_layer []*Knot
	for i:=0; i<len(layer); i++ {
			var kn1, kn2 *Knot
			s, kn1, kn2 = layer[i].AddTwoKnots(s)
			if kn1 != nil {
				new_layer = append(new_layer, kn1)
			}	
			if kn2 != nil {
				new_layer = append(new_layer, kn2)
			}
		}
		
	return s, new_layer
}


func (t *tree) String() string {
	var s string = fmt.Sprint(t.root.value)
	var aktuelle_layer []*Knot
	aktuelle_layer = append(aktuelle_layer, t.root)

	fmt.Println(t.root.leftPtr, t.root.rightPtr)
	
	for {
		s = s + "\n"
		s, aktuelle_layer = interate_trough_layer(s, aktuelle_layer)
		if len(aktuelle_layer) == 0 {
			break
		}
	}
	
	return s
}

		
