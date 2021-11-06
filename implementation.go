package lab2

import (
	"errors"
	"strings"
)

func PostfixToPrefix(input string) (res string, err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("invalid syntax")
		}
	}()
	res = strings.TrimSuffix(PrintPrefix("", GetTree(input)), " ")
	return res, err
}

type Node struct {
	Symbol string
	Left   *Node
	Right  *Node
	Parent *Node
}

func NewNode(symbol string) Node {
	m := Node{
		Symbol: symbol,
	}
	return m
}

func (node *Node) IsOperator() bool {
	switch node.Symbol {
	case
		"+",
		"-",
		"*",
		"/",
		"^":
		return true
	}
	return false
}

func (parent *Node) AddChildToTree(child *Node) {
	if parent.Right == nil {
		parent.AddRightChild(child)
	} else if parent.Left == nil {
		parent.AddLeftChild(child)
	} else {
		parent.Parent.AddChildToTree(child)
	}
}

func (parent *Node) AddRightChild(child *Node) {
	parent.Right = child
	child.Parent = parent
}

func (parent *Node) AddLeftChild(child *Node) {
	parent.Left = child
	child.Parent = parent
}

func PrintPrefix(print string, node *Node) string {
	print += node.Symbol + " "
	if node.Left != nil {
		print = PrintPrefix(print, node.Left)
		print = PrintPrefix(print, node.Right)
	}
	return print
}

func GetTree(input string) *Node {
	slice := Revert(strings.Split(input, " "))
	var root, parent *Node

	for _, symbol := range slice {
		current := NewNode(symbol)

		if root == nil {
			root = &current
		} else {
			parent.AddChildToTree(&current)
		}

		if current.IsOperator() {
			parent = &current
		}
	}

	return root
}

func Revert(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
