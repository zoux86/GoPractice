package main

import (
	"container/list"
	"fmt"
	)

func main(){
	l := list.New()
	l.PushBack(1)
	l.PushBack("hh")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	var s *list.List
	s = list.New()
	s.PushFront(1)
	s.PushFront("aa")
	for e := s.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}


}

