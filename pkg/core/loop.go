package core

type CircularList struct {
	root *Element
	cur  *Element
}

func NewList() *CircularList {
	root := &Element{}
	root.next = root
	root.prev = root
	return &CircularList{root: root, cur: root}
}

func (c *CircularList) Add(v *Value) *Element {
	e := &Element{
		V:    v,
		next: c.root.next,
		prev: c.root,
	}

	c.root.next.prev = e
	c.root.next = e

	return e
}

func (c *CircularList) Del(e *Element) {
	e.prev.next = e.next
	e.next.prev = e.prev
}

func (c *CircularList) Next() *Element {
	c.cur = c.cur.next

	if c.cur == c.root {
		return c.Next()
	}

	return c.cur
}

type Element struct {
	V *Value

	next *Element
	prev *Element
}

type Value struct {
}
