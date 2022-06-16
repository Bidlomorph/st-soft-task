package main

import (
	"fmt"
)

type element struct {
	value string
	index uint
}

type listElement struct {
	element element
	next    *listElement
}

type list struct {
	len  int
	head *listElement
}

func initList() *list {
	return &list{}
}

func (s *list) InsertElementBefore(
	prev *listElement,
	current element) {
	listElement := &listElement{
		element: current,
	}

	listElement.next = prev
	s.head = listElement
	return
}

func (s *list) InsertElementAfter(
	prev, next *listElement,
	current element) {
	listElement := &listElement{
		element: current,
	}

	if next == nil {
		prev.next = listElement
	} else {
		prev.next = listElement
		listElement.next = next
	}
	return
}

func (s *list) Insert(elem element) (err error) {
	if s.head == nil {
		s.InsertElementBefore(nil, elem)
		return
	}

	for sc := s.head; sc != nil; sc = sc.next {
		prev := sc
		next := sc.next

		if prev.element.index == elem.index {
			return fmt.Errorf("element: %v - alredy in list", elem)
		}

		if prev.element.index > elem.index {
			s.InsertElementBefore(prev, elem)
			s.len++
			break
		}

		if next == nil {
			s.InsertElementAfter(prev, next, elem)
			s.len++
			break
		}

		if prev.element.index < elem.index && next.element.index > elem.index {
			if next.element.index == elem.index {
				return fmt.Errorf("element: %v - alredy in list", elem)
			}
			s.InsertElementAfter(prev, next, elem)
			s.len++
			break
		}
	}
	return
}

func (s *list) Show() {
	for sc := s.head; sc != nil; sc = sc.next {
		fmt.Println(sc.element)
	}
}

func main() {
	list := initList()
	// Initialize list
	if err := list.Insert(element{"A", 1}); err != nil {
		fmt.Println(err)
	}
	if err := list.Insert(element{"B", 4}); err != nil {
		fmt.Println(err)
	}
	if err := list.Insert(element{"C", 7}); err != nil {
		fmt.Println(err)
	}
	if err := list.Insert(element{"D", 12}); err != nil {
		fmt.Println(err)
	}
	if err := list.Insert(element{"E", 26}); err != nil {
		fmt.Println(err)
	}

	// Add some more indexes
	if err := list.Insert(element{"B1", 6}); err != nil {
		fmt.Println(err)
	}
	if err := list.Insert(element{"E1", 27}); err != nil {
		fmt.Println(err)
	}
	if err := list.Insert(element{"C1", 8}); err != nil {
		fmt.Println(err)
	}
	if err := list.Insert(element{"D", 12}); err != nil {
		fmt.Println(err)
	}
	if err := list.Insert(element{"A0", 0}); err != nil {
		fmt.Println(err)
	}

	list.Show()
}
