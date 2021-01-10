package ntrees

import (
	"tree/utils"
)

type NNode struct {
	Value int
	Children []*NNode
}

func (n *NNode) childrenCount() int {
	return len(n.Children)
}

func NTree(value int, children []*NNode) *NNode {
	n := &NNode{
		Value:    value,
		Children: children,
	}
	return n
}

func (n *NNode) Op() {
	ch := make(chan utils.Nothing, n.childrenCount())
	sum := 0

	for _, node := range n.Children {
		sum += node.Value
		node := node
		go func() {
			node.Op()
			ch <- utils.Nothing{}
		}()
	}

	n.Value = sum
	for i := 0; i < n.childrenCount(); i++ {
		<-ch
	}
}
