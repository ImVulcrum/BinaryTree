package binarytree

import (
	"fmt"
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

func (current_knot *knot) add_childs_of_knot_to_layer(new_layer []*knot) []*knot {
	if current_knot == nil {
		new_layer = append(new_layer, nil, nil)
	} else {
		new_layer = append(new_layer, current_knot.leftPtr)
		new_layer = append(new_layer, current_knot.rightPtr)
	}
	return new_layer
}

func interate_trough_layer(layer []*knot) []*knot {
	var new_layer []*knot

	for i := 0; i < len(layer); i++ {
		new_layer = layer[i].add_childs_of_knot_to_layer(new_layer)
	}
	return new_layer
}

func (t *tree) Print(terminal_width int) string {
	var s string
	var current_layer []*knot

	for layer_count := 0; true; layer_count++ {
		if layer_count == 0 {
			current_layer = append(current_layer, t.root)
		} else {
			current_layer = interate_trough_layer(current_layer)
		}

		if layer_empty(current_layer) {
			break
		}

		fmt.Println(" " + create_layer_string(current_layer, t.height-layer_count, t.max_len_of_value_string))
	}

	return s
}

func create_layer_string(layer []*knot, reverse_layer_count int, max_count_of_value_digits int) string {
	var current_element_equalized string
	var layer_string string
	var spacing_count int = calculate_layer_spacing(reverse_layer_count, max_count_of_value_digits)
	var modified_value_digit_count int = max_count_of_value_digits

	if spacing_count == 0 { //for some reason this has to be set to zero if it is the second last layer
		modified_value_digit_count = 0
	}

	for i := 0; i < len(layer); i++ {
		current_element_equalized = layer[i].get_value_string_equalized(max_count_of_value_digits)

		if reverse_layer_count != 1 { //if its is not the last layer
			if i == 0 { //this spacing is only neaded on the leftest element
				layer_string = strings.Repeat(" ", max_count_of_value_digits/2)
			}

			if layer[i] != nil {
				//add spacing to the left
				layer_string = layer_string + strings.Repeat(" ", spacing_count+modified_value_digit_count/2)

				if layer[i].leftPtr != nil {
					layer_string = layer_string + "┌" + strings.Repeat("─", modified_value_digit_count/2+spacing_count) + current_element_equalized
				} else {
					layer_string = layer_string + " " + strings.Repeat(" ", modified_value_digit_count/2+spacing_count) + current_element_equalized
				}

				if layer[i].rightPtr != nil {
					layer_string = layer_string + strings.Repeat("─", modified_value_digit_count/2+spacing_count) + "┐"
				} else {
					layer_string = layer_string + strings.Repeat(" ", modified_value_digit_count/2+spacing_count) + " "
				}

				//add spacing to the right
				layer_string = layer_string + strings.Repeat(" ", modified_value_digit_count/2+spacing_count+max_count_of_value_digits)

			} else { //if the current element is nil, treat it like an element with no childs
				//left spacing + leftPtr (else block) + rightPtr (else block) + right spacing
				layer_string = layer_string + strings.Repeat(" ", 4*spacing_count+2*modified_value_digit_count+2*max_count_of_value_digits+2)
			}

		} else { //if it is the last layer
			layer_string = layer_string + current_element_equalized + " "
		}
	}

	return layer_string
}

func calculate_layer_spacing(reverse_layer_count int, max_len_of_value_string int) int {
	var iteration_count int = reverse_layer_count - 3
	var layer_spacing int = 1

	if iteration_count == -1 {
		return 0
	} else if iteration_count == -2 {
		return -2
	}

	for i := 0; i < iteration_count; i++ {
		layer_spacing = (2 * layer_spacing) + max_len_of_value_string
	}

	return layer_spacing
}

func (k *knot) get_value_string() string {
	if k != nil {
		if len(fmt.Sprint(k.value))%2 == 0 {
			return "0" + fmt.Sprint(k.value)
		}
		return fmt.Sprint(k.value)
	}
	return " "
}

func (k *knot) get_value_string_equalized(max_len_of_value_string int) string {
	var value_string_unequlized string = k.get_value_string()

	var has_left_child bool
	var has_right_child bool

	if k != nil {
		has_left_child = k.leftPtr != nil
		has_right_child = k.rightPtr != nil
	}

	return equalize_different_value_lenghts(value_string_unequlized, max_len_of_value_string, has_left_child, has_right_child)
}

func layer_empty(layer []*knot) bool {
	for i := 0; i < len(layer); i++ {
		if layer[i] != nil {
			return false
		}
	}
	return true
}

func equalize_different_value_lenghts(current_value_string string, max_string_size int, has_left_child bool, has_right_child bool) string {
	var line_seperator string = strings.Repeat("─", (max_string_size-len(current_value_string))/2)
	var space_seperator string = strings.Repeat(" ", (max_string_size-len(current_value_string))/2)

	if has_left_child && has_right_child {
		return line_seperator + current_value_string + line_seperator
	} else if has_left_child && !has_right_child {
		return line_seperator + current_value_string + space_seperator
	} else if !has_left_child && has_right_child {
		return space_seperator + current_value_string + line_seperator
	} else { //!has_left_child && !has_right_child
		return space_seperator + current_value_string + space_seperator
	}
}

// func calc_len_of_bottom_layer_string(tree_height int, max_len_of_value_string int) int {
// 	if tree_height == 1 {
// 		return 1
// 	}
// 	return max_len_of_value_string*int(math.Pow(2, float64(tree_height-1))) + int(math.Pow(2, float64(tree_height-1))) + 1
// }

// func calc_front_spacing(tree_height int, current_string string, max_len_of_value_string int) int {
// 	var spacing int = (calc_len_of_bottom_layer_string(tree_height, max_len_of_value_string) - len(current_string)) / 2
// 	if spacing < 0 {
// 		return 0
// 	}
// 	return spacing
// }

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
