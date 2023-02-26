package core

import (
	"testing"
)

type MockSentient struct {
}

func (m MockSentient) AddEnergy(energy int) {
	panic("implement me")
}

func (m MockSentient) SubEnergy(energy int) {
	//TODO implement me
	panic("implement me")
}

func (m MockSentient) GetEnergy() int {
	panic("implement me")
}

func (m MockSentient) Act() int {
	panic("implement me")
}

func TestLoop_Add(t *testing.T) {
	list := NewLoop()
	e1 := list.Add(&MockSentient{})
	e2 := list.Add(&MockSentient{})
	e3 := list.Add(&MockSentient{})
	//root->e3->e2->e1->root

	if e3 != list.Next() {
		t.Fatalf("3: the loop is broken")
	}
	if e2 != list.Next() {
		t.Fatalf("2: the loop is broken")
	}
	if e1 != list.Next() {
		t.Fatalf("1: the loop is broken")
	}
	if nil != list.Next() {
		t.Fatalf("3/loop again: the loop is broken")
	}
}

func TestLoop_Del(t *testing.T) {
	list := NewLoop()
	e1 := list.Add(&MockSentient{})
	e2 := list.Add(&MockSentient{})
	e3 := list.Add(&MockSentient{})
	//root->e3->e2->e1->root

	list.Del(e2)
	//root->e3->e1->root

	if e3 != list.Next() {
		t.Fatalf("3: the loop is broken")
	}

	if e1 != list.Next() {
		t.Fatalf("1: the loop is broken")
	}

	if nil != list.Next() {
		t.Fatalf("3/loop again: the loop is broken")
	}
}
