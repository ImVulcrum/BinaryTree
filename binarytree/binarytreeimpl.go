package binarytree

import (
	"fmt"
	"strings"
)

type tree struct {
	knots                   []*knot
	root                    *knot
	current_knot            *knot
	height                  int
	max_len_of_value_string int
	is_avl_tree             bool
}

type knot struct {
	leftPtr        *knot
	rightPtr       *knot
	value          int
	balance_factor int
}

func NewKnot(value int) *knot {
	var k *knot = new(knot)
	k.value = value
	k.leftPtr = nil
	k.rightPtr = nil
	return k
}

func NewTree(is_avl_tree bool) *tree {
	var t *tree = new(tree)
	t.is_avl_tree = is_avl_tree
	return t
}

func (t *tree) GiveKnots() []*knot {
	return t.knots
}

func (t *tree) Search(value int) *knot {
	return t.search(t.root, value)
}

func (t *tree) String() string {
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
		s = s + " " + create_layer_string(current_layer, t.height-layer_count, t.max_len_of_value_string) + "\n"
	}
	return s
}

func (t *tree) InsertList(values []int) {
	for _, v := range values {
		t.Insert(v)
	}
}

func (t *tree) Insert(value int) {
	var knot_to_be_inserted *knot = NewKnot(value)
	t.knots = append(t.knots, knot_to_be_inserted)

	if len(fmt.Sprint(value)) > t.max_len_of_value_string {
		t.max_len_of_value_string = len(fmt.Sprint(value))
		if t.max_len_of_value_string%2 == 0 {
			t.max_len_of_value_string++
		}
	}

	if t.root == nil {
		t.root = knot_to_be_inserted
		t.height = 1
		return
	}

	t.current_knot = t.root
	for {
		if value <= t.current_knot.value && t.current_knot.leftPtr != nil {
			t.current_knot = t.current_knot.leftPtr
		} else if value > t.current_knot.value && t.current_knot.rightPtr != nil {
			t.current_knot = t.current_knot.rightPtr
		} else {
			break
		}
	}

	if value <= t.current_knot.value {
		t.current_knot.leftPtr = knot_to_be_inserted
	} else if value > t.current_knot.value {
		t.current_knot.rightPtr = knot_to_be_inserted
	}

	if t.is_avl_tree {
		var problem_knot *knot = t.calculate_balance_factors()
		if problem_knot != nil { //rotate
			t.rotate(problem_knot)
		}
	}
	t.height = t.root.calculate_height_from_bottom()
}

func (t *tree) rotate(problem_knot *knot) {
	root_of_problem_knot, problem_knot_is_left_child_of_root := t.get_previous_knot(t.root, problem_knot)

	if problem_knot.balance_factor == 2 {
		if problem_knot.rightPtr.balance_factor >= 0 { //simple left rotation
			t.left_rotation(root_of_problem_knot, problem_knot, problem_knot_is_left_child_of_root)
		} else { //right-left rotation
			t.right_rotation(problem_knot, problem_knot.rightPtr, false)
			t.left_rotation(root_of_problem_knot, problem_knot, problem_knot_is_left_child_of_root)
		}
	} else if problem_knot.balance_factor == -2 {
		if problem_knot.leftPtr.balance_factor <= 0 { //simple right rotation
			t.right_rotation(root_of_problem_knot, problem_knot, problem_knot_is_left_child_of_root)
		} else { //left-right rotation
			t.left_rotation(problem_knot, problem_knot.leftPtr, true)
			t.right_rotation(root_of_problem_knot, problem_knot, problem_knot_is_left_child_of_root)
		}
	}
}

func (t *tree) left_rotation(root_of_problem_knot *knot, problem_knot *knot, problem_knot_is_on_the_left_of_root bool) {
	var lower_knot *knot          //right_knot with balance factor 1 or 0
	var subtree_to_be_moved *knot //middle knot

	//save the knots
	lower_knot = problem_knot.rightPtr
	subtree_to_be_moved = lower_knot.leftPtr

	//move the middle knot the the problem knot (with balance factor two)
	problem_knot.rightPtr = subtree_to_be_moved

	//add the problem knot to the left side of the lower knot
	lower_knot.leftPtr = problem_knot

	//add the lower knot the root
	if root_of_problem_knot != nil { //the problem knot is not the root of the tree
		if problem_knot_is_on_the_left_of_root {
			root_of_problem_knot.leftPtr = lower_knot
		} else {
			root_of_problem_knot.rightPtr = lower_knot
		}
	} else { //the problem knot is the root of the tree
		t.root = lower_knot
	}
}

func (t *tree) right_rotation(root_of_problem_knot *knot, problem_knot *knot, problem_knot_is_on_the_left_of_root bool) {
	var lower_knot *knot          //right_knot with balance factor -1 or 0
	var subtree_to_be_moved *knot //middle knot

	//save the knots
	lower_knot = problem_knot.leftPtr
	subtree_to_be_moved = lower_knot.rightPtr

	//move the middle knot the the problem knot (with balance factor minus two)
	problem_knot.leftPtr = subtree_to_be_moved

	//add the problem knot to the left right of the lower knot
	lower_knot.rightPtr = problem_knot

	//add the lower knot the root
	if root_of_problem_knot != nil { //the problem knot is not the root of the tree
		if problem_knot_is_on_the_left_of_root {
			root_of_problem_knot.leftPtr = lower_knot
		} else {
			root_of_problem_knot.rightPtr = lower_knot
		}
	} else { //the problem knot is the root of the tree
		t.root = lower_knot
	}
}

func (t *tree) get_previous_knot(relative_root *knot, knot_to_find_parent *knot) (*knot, bool) { //if the knot is the left child of the parent -> true; if it's the right -> false
	if relative_root != nil {
		if knot_to_find_parent.value < relative_root.value {
			if relative_root.leftPtr == knot_to_find_parent {
				return relative_root, true
			} else {
				return t.get_previous_knot(relative_root.leftPtr, knot_to_find_parent)
			}
		} else if knot_to_find_parent.value > relative_root.value {
			if relative_root.rightPtr == knot_to_find_parent {
				return relative_root, false
			} else {
				return t.get_previous_knot(relative_root.rightPtr, knot_to_find_parent)
			}
		} else { //this knot has no previous knot, therefore it's the root knot
			return nil, false
		}
	} else {
		panic("this function is broken")
	}
}

func (t *tree) search(relative_root *knot, value int) *knot {
	if relative_root == nil {
		return nil
	} else if value < relative_root.value {
		return t.search(relative_root.leftPtr, value)
	} else if value > relative_root.value {
		return t.search(relative_root.rightPtr, value)
	} else {
		return relative_root
	}
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
				layer_string = layer_string + strings.Repeat(" ", 4*spacing_count+4*(modified_value_digit_count/2)+2+2*max_count_of_value_digits)
				// layer_string = layer_string + strings.Repeat(" ", spacing_count+modified_value_digit_count/2) + " " + strings.Repeat(" ", modified_value_digit_count/2+spacing_count) +
				// current_element_equalized + strings.Repeat(" ", modified_value_digit_count/2+spacing_count) + " " + strings.Repeat(" ", modified_value_digit_count/2+spacing_count+max_count_of_value_digits)
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

func (t *tree) calculate_balance_factors() *knot {
	var lowest_problem_knot *knot
	var current_knot *knot

	for i := 0; i < len(t.knots); i++ {
		current_knot = t.knots[i]
		current_knot.calculate_balance_factor()
		if current_knot.balance_factor == 2 || current_knot.balance_factor == -2 {
			if (lowest_problem_knot != nil && current_knot.calculate_distance_to_root(t.height) > lowest_problem_knot.calculate_distance_to_root(t.height)) || lowest_problem_knot == nil {
				lowest_problem_knot = current_knot
			}
		}
	}
	return lowest_problem_knot
}

func (k *knot) calculate_balance_factor() int {
	if k == nil {
		k.balance_factor = 0
		return 0
	} else {
		var left_height int = k.leftPtr.calculate_height_from_bottom()
		var right_height int = k.rightPtr.calculate_height_from_bottom()
		k.balance_factor = right_height - left_height
		return right_height - left_height
	}
}

func (k *knot) calculate_distance_to_root(tree_height int) int {
	return tree_height - k.calculate_height_from_bottom()
}

func (k *knot) calculate_height_from_bottom() int {
	if k == nil {
		return 0
	} else {
		var left_height int = k.leftPtr.calculate_height_from_bottom()
		var right_height int = k.rightPtr.calculate_height_from_bottom()

		if left_height > right_height {
			return left_height + 1
		} else {
			return right_height + 1
		}
	}
}
