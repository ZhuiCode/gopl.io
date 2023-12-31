package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	val := appendValues(values[:0], root)
	fmt.Println(val)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	values := []int{2, 3, 4, 4, 5, 2, 5, 19, 3, 32, 44}
	sort(values)
}
