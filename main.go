package main

import "fmt"



type Node struct {
	value interface{}
	next *Node
}

type Stack struct {
	top *Node
	length int
}

func NewStack() *Stack {
	stack := &Stack{
		top: nil,
		length: 0,
	}
	return stack
}

func (s *Stack) Empty() bool {
	return s.top == nil
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Top() interface{} {
	return s.top
}

func (s *Stack) Push(value interface{}) {
	node := &Node{
		value: value,
		next: s.top,
	}
	s.top = node
}

func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	item := s.top
	s.top = s.top.next
	return item
}

func main() {

	stack := NewStack()
	//stack.Push(12)
	//stack.Push(23)

	stack.Push(34)
	fmt.Println(stack.Pop().(*Node).value)
	//fmt.Println(stack.Pop().(*Node).value)
	//fmt.Println(stack.Pop().(*Node).value)
}



////import "fmt"
//
//var top *node
//type node struct {
//	data string
//	next *node
//}
//
//func (st node) push(data string){
//	pr := new (node)
//	pr.next = st.top()
//}
//func(st node) pop(){
//
//}
//func (st node) top()(topdata string){
//
//}
//func (st node) empty(){
//
//}
//func (st node) init(){
//
//}
//type stack interface{
//	push(data string)
//	pop()
//	top()
//	empty()
//	init()
//}
//func main(){
//
//}