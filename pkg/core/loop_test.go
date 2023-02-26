package core

import (
	"testing"
)

func TestCircularList_Add(t *testing.T) {
	list := NewList()
	e1 := list.Add(&Value{})
	e2 := list.Add(&Value{})
	e3 := list.Add(&Value{})
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
	if e3 != list.Next() {
		t.Fatalf("3/loop again: the loop is broken")
	}
}

func TestCircularList_Del(t *testing.T) {
	list := NewList()
	e1 := list.Add(&Value{})
	e2 := list.Add(&Value{})
	e3 := list.Add(&Value{})
	//root->e3->e2->e1->root

	list.Del(e2)
	//root->e3->e1->root

	if e3 != list.Next() {
		t.Fatalf("3: the loop is broken")
	}

	if e1 != list.Next() {
		t.Fatalf("1: the loop is broken")
	}

	if e3 != list.Next() {
		t.Fatalf("3/loop again: the loop is broken")
	}
}
