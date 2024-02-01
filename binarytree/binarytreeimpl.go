package binarytree

import (
	"fmt"
	"strings"
)

type tree struct {
	knots        []*knot
	root         *knot
	current_knot *knot
	knot_count   int
	height       int
}

type knot struct {
	leftPtr  *knot
	rightPtr *knot
	value    int
}

func NewKnot(value int) *knot {
	var k *knot = new(knot)
	k.value = value
	k.leftPtr = nil
	k.rightPtr = nil
	return k
}

func NewTree() *tree {
	var t *tree = new(tree)
	return t
}

func (t *tree) GiveKnots() []*knot {
	return t.knots
}

func (t *tree) InsertList(values []int) {
	for _, v := range values {
		t.Insert(v)
	}
}

func (t *tree) Insert(value int) {
	var knot *knot = NewKnot(value)
	t.knot_count++
	t.knots = append(t.knots, knot)

	if t.root == nil {
		t.root = knot
		t.height = 1
		return
	}

	t.current_knot = t.root
	count := 2
	for {
		if value <= t.current_knot.value && t.current_knot.leftPtr != nil {
			t.current_knot = t.current_knot.leftPtr
			count++
		} else if value > t.current_knot.value && t.current_knot.rightPtr != nil {
			t.current_knot = t.current_knot.rightPtr
			count++
		} else {
			break
		}
	}
	if count > t.height {
		t.height = count
	}

	if value <= t.current_knot.value {
		t.current_knot.leftPtr = knot
	} else if value > t.current_knot.value {
		t.current_knot.rightPtr = knot
	}
}

func (t *tree) search(value int) *knot {
	if t.current_knot == nil {
		return nil
	} else if value < t.current_knot.value {
		t.current_knot = t.current_knot.leftPtr
		return t.search(value)
	} else if value > t.current_knot.value {
		t.current_knot = t.current_knot.rightPtr
		return t.search(value)
	} else {
		return t.current_knot
	}
}

func (t *tree) Search(value int) *knot {
	t.current_knot = t.root
	knot := t.search(value)
	return knot
}

func (current_knot *knot) add_childs_of_knot_to_string(new_layer []*knot, layer_spacing int) (string, []*knot) {
	var left_knot, right_knot string = "n", "n"
	var string_addition string

	if current_knot == nil {
		new_layer = append(new_layer, nil, nil)
		left_knot = " "
		right_knot = " "

	} else {

		if current_knot.leftPtr != nil {
			left_knot = fmt.Sprint(current_knot.leftPtr.value)
		} else {
			left_knot = " "
		}

		new_layer = append(new_layer, current_knot.leftPtr)

		if current_knot.rightPtr != nil {
			right_knot = fmt.Sprint(current_knot.rightPtr.value)

		} else {
			right_knot = " "
		}

		new_layer = append(new_layer, current_knot.rightPtr)

	}

	var separator string = strings.Repeat(" ", layer_spacing)

	string_addition = left_knot + separator + right_knot

	return string_addition, new_layer
}

func interate_trough_layer(layer []*knot, layer_spacing int) (string, []*knot) {
	var new_layer []*knot
	var layer_string string
	var layer_addition_string string
	var separator string = strings.Repeat(" ", layer_spacing)

	for i := 0; i < len(layer); i++ {
		layer_addition_string, new_layer = layer[i].add_childs_of_knot_to_string(new_layer, layer_spacing)
		if layer_string != "" {
			layer_string = layer_string + separator + layer_addition_string
		} else {
			layer_string = layer_string + layer_addition_string
		}
	}
	return layer_string, new_layer
}

func (t *tree) Print(terminal_width int) string {
	var s string
	var string_addition string = fmt.Sprint(t.root.value)

	var current_layer []*knot
	current_layer = append(current_layer, t.root)

	var layer_count int = 1

	for {
		fmt.Println(centerString(string_addition, terminal_width))

		layer_count++
		s = s + string_addition + "\n"
		string_addition, current_layer = interate_trough_layer(current_layer, calculate_layer_spacing(t.height, layer_count))
		if !layer_not_empty(current_layer) {
			break
		}
	}

	return s
}

func calculate_layer_spacing(tree_height, layer_count int) int {
	var iteration_count int = tree_height - layer_count
	var layer_spacing int = 1

	for i := 0; i < iteration_count; i++ {
		layer_spacing = (2 * layer_spacing) + 1
	}
	return layer_spacing
}

func layer_not_empty(layer []*knot) bool {
	for i := 0; i < len(layer); i++ {
		if layer[i] != nil {
			return true
		}
	}
	return false
}

func centerString(s string, width int) string {
	sLen := len(s)
	if sLen >= width {
		return s
	}

	// Calculate the number of spaces needed on each side
	padding := (width - sLen) / 2

	// Build the centered string
	centeredStr := strings.Repeat(" ", padding) + s + strings.Repeat(" ", padding)

	return centeredStr
}
