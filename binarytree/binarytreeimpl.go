package binarytree

import (
	"fmt"
	"math"
	"strings"
)

type tree struct {
	knots                   []*knot
	root                    *knot
	current_knot            *knot
	knot_count              int
	height                  int
	max_len_of_value_string int
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

	if len(fmt.Sprint(value)) > t.max_len_of_value_string {
		t.max_len_of_value_string = len(fmt.Sprint(value))
		if t.max_len_of_value_string%2 == 0 {
			t.max_len_of_value_string++
		}
	}

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

func (current_knot *knot) add_childs_of_knot_to_string(new_layer []*knot, layer_spacing int, max_len int) (string, []*knot) {
	var left_knot, right_knot string = "n", "n"
	var string_addition string
	var separator string = strings.Repeat(" ", layer_spacing)

	if current_knot == nil {
		new_layer = append(new_layer, nil, nil)
		left_knot = "n"
		right_knot = "n"

	} else {

		if current_knot.leftPtr != nil {
			left_knot = fmt.Sprint(current_knot.leftPtr.value)
		} else {
			left_knot = "n"
		}

		new_layer = append(new_layer, current_knot.leftPtr)

		if current_knot.rightPtr != nil {
			right_knot = fmt.Sprint(current_knot.rightPtr.value)

		} else {
			right_knot = "n"
		}

		new_layer = append(new_layer, current_knot.rightPtr)
	}

	if len(left_knot)%2 == 0 {
		left_knot = "0" + left_knot
	}
	if len(right_knot)%2 == 0 {
		right_knot = "0" + right_knot
	}

	if len(left_knot) < max_len {
		left_knot = equalize_different_value_lenghts(left_knot, max_len)
	}
	if len(right_knot) < max_len {
		right_knot = equalize_different_value_lenghts(right_knot, max_len)
	}

	string_addition = left_knot + separator + right_knot

	return string_addition, new_layer
}

func interate_trough_layer(layer []*knot, layer_spacing int, max_len_of_value_string int) (string, []*knot) {
	var new_layer []*knot
	var layer_string string
	var layer_addition_string string
	var separator string = strings.Repeat(" ", layer_spacing)

	for i := 0; i < len(layer); i++ {
		layer_addition_string, new_layer = layer[i].add_childs_of_knot_to_string(new_layer, layer_spacing, max_len_of_value_string)
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
	var string_addition string
	var current_layer []*knot

	for layer_count := 0; true; layer_count++ {
		if layer_count == 0 {
			current_layer = append(current_layer, t.root)
			string_addition = fmt.Sprint(t.root.value)
			string_addition = equalize_different_value_lenghts(string_addition, t.max_len_of_value_string)
		} else {
			string_addition, current_layer = interate_trough_layer(current_layer, calculate_layer_spacing(t.height, layer_count+1, t.max_len_of_value_string), t.max_len_of_value_string)
		}

		if layer_empty(current_layer) {
			break
		}

		string_addition = strings.Repeat(" ", calc_front_spacing(t.height, string_addition, t.max_len_of_value_string)) + string_addition
		s = s + string_addition + "\n"

		fmt.Println(string_addition)
	}

	return s
}

func calculate_layer_spacing(tree_height, layer_count int, max_len_of_value_string int) int {
	var iteration_count int = tree_height - layer_count
	var layer_spacing int = 1

	for i := 0; i < iteration_count; i++ {
		layer_spacing = (2 * layer_spacing) + max_len_of_value_string
	}

	return layer_spacing
}

func layer_empty(layer []*knot) bool {
	for i := 0; i < len(layer); i++ {
		if layer[i] != nil {
			return false
		}
	}
	return true
}

func calc_len_of_bottom_layer_string(tree_height int, max_len_of_value_string int) int {
	if tree_height == 1 {
		return 1
	}
	return max_len_of_value_string*int(math.Pow(2, float64(tree_height-1))) + int(math.Pow(2, float64(tree_height-1))) + 1
}

func calc_front_spacing(tree_height int, current_string string, max_len_of_value_string int) int {
	var spacing int = (calc_len_of_bottom_layer_string(tree_height, max_len_of_value_string) - len(current_string)) / 2
	if spacing < 0 {
		return 0
	}
	return spacing
}

func equalize_different_value_lenghts(current_value_string string, max_string_size int) string {
	var seperator string = strings.Repeat("-", (max_string_size-len(current_value_string))/2)
	return seperator + current_value_string + seperator
}

// func centerString(s string, width int) string {
// 	sLen := len(s)
// 	if sLen >= width {
// 		return s
// 	}

// 	// Calculate the number of spaces needed on each side
// 	padding := (width - sLen) / 2

// 	// Build the centered string
// 	centeredStr := strings.Repeat(" ", padding) + s + strings.Repeat(" ", padding)

// 	return centeredStr
// }
