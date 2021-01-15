package ttrees

type TNode struct {
	Value int
	Children [3]*TNode
}

func TTree(value int, children [3]*TNode) *TNode {
	n := &TNode{
		Value:    value,
		Children: children,
	}
	return n
}

func (n *TNode) Op() {
	max := n.Value
	count := 0
	ch := make(chan int, 3)
	for _, c := range n.Children {
		if c == nil {
			continue
		}
		count++
		c := c
		go func() {
			c.Op()
			ch <- c.Value
		}()
	}

	for i := 0; i < count; i++ {
		v := <- ch
		if v > max {
			max = v
		}
	}

	n.Value = max
}
