package core

type Loop struct {
	root *Element
	cur  *Element
}

func NewLoop() *Loop {
	root := &Element{}
	root.next = root
	root.prev = root
	return &Loop{root: root, cur: root}
}

func (l *Loop) Add(s Sentient) *Element {
	e := &Element{
		S:    s,
		next: l.root.next,
		prev: l.root,
	}

	l.root.next.prev = e
	l.root.next = e

	return e
}

func (l *Loop) Del(e *Element) {
	e.prev.next = e.next
	e.next.prev = e.prev
}

func (l *Loop) Next() *Element {
	l.cur = l.cur.next

	if l.cur == l.root {
		return nil
	}

	return l.cur
}

type Element struct {
	S Sentient

	next *Element
	prev *Element
}

type Sentient interface {
	AddEnergy(energy int)
	SubEnergy(energy int)
	GetEnergy() int
	Act() int
}
