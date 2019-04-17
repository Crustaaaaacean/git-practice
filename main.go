package main

import (
	"fmt"
	"strconv"
	"unicode"
)

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
	return s.top.value
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
	item := s.top.value
	s.top = s.top.next
	return item
}

func inorderToPostorder(inorder string) string{
	var inex = []byte(inorder)
	var lenex = len(inorder)
	var postorder string
	var tmp byte
	var i int
	stack := NewStack()
	for i = 0; i < lenex;{
		switch inex[i] {
		case ' ':
			i++
			continue
		case '(':
			stack.Push(inex[i])
			i++
		case ')':
			tmp = stack.Pop().(byte)
			for tmp != '(' {
				postorder += string(tmp)
				tmp = stack.Pop().(byte)
			}
			i++
		case '*','/':
			for stack.Empty() != true {
				tmp = stack.Top().(byte)
				if tmp == '*' || tmp == '/' {
					postorder += string(tmp)
					tmp = stack.Pop().(byte)
				} else {
					break
				}
			}
			stack.Push(inex[i])
			i++
		case '+','-':
			for stack.Empty() != true {
				tmp = stack.Top().(byte)
				if tmp != '(' {
					postorder += string(tmp)
					tmp = stack.Pop().(byte)
				} else {
					break
				}
			}
			stack.Push(inex[i])
			i++
		default:
			for inex[i] >= '0' && inex[i] <= '9' {
				postorder += string(inex[i])
				i++
				if i >= lenex {
					break
				}
			}
			postorder += "#"
		}
	}
	for stack.Empty() != true {
		tmp = stack.Pop().(byte)
		postorder += string(tmp)
	}
	return postorder
}

func calPostorder(postorder string) int{
	var res int
	var tmp string
	var tmpi int
	var n1, n2 int
	stack := NewStack()
	for i := 0; i < len(postorder); i++ {
		tmp = ""
		if unicode.IsDigit(rune(postorder[i])) {
			for postorder[i] != '#' {
				tmp += string(postorder[i])
				i++
			}
			tmpi, _ = strconv.Atoi(tmp)
			stack.Push(tmpi)
		} else {
			n1 = stack.Pop().(int)
			n2 = stack.Pop().(int)
			switch postorder[i] {
			case '+':
				stack.Push(n2 + n1)
			case '-':
				stack.Push(n2 - n1)
			case '*':
				stack.Push(n2 * n1)
			case '/':
				stack.Push(n2 / n1)//注意这块的计算n2在前
			}
		}
	}
	res = stack.Pop().(int)
	return res
}

func main(){
	//var inorder string = "12+31*5/6"
	var inorder string
	var	postorder string
	fmt.Println("input inorder expression")
	fmt.Scanf("%s", &inorder)
	fmt.Printf("输入的中缀表达式:%s\n", inorder)
	postorder = inorderToPostorder(inorder)
	fmt.Printf("后缀表达式:%s\n", postorder)
	fmt.Printf("计算结果是:%d",calPostorder(postorder))
}
