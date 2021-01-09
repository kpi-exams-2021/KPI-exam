package trees

type nothing struct {}

type Node struct {
	ch    chan nothing
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) hasLeft() bool {
	return n.Left != nil
}

func (n *Node) hasRight() bool {
	return n.Right != nil
}

func Tree(value int, left *Node, right *Node) *Node {
	return &Node{
		ch:    make(chan nothing, 2),
		Value: value,
		Left:  left,
		Right: right,
	}
}

func (n *Node) Sum() {
	sum := 0
	wait := 0

	if n.hasLeft() {
		sum += n.Left.Value
		wait++
		go func() {
			n.Left.Sum()
			n.ch <- nothing{}
		}()
	}

	if n.hasRight() {
		sum += n.Right.Value
		wait++
		go func() {
			n.Right.Sum()
			n.ch <- nothing{}
		}()
	}

	n.Value = sum
	for i := 0; i < wait; i++ {
		<- n.ch
	}
}

func (n *Node) Init() {
	n.ch = make(chan nothing, 2)
}

func (n *Node) ForEach(operation func(n *Node)) {
	operation(n)
	if n.hasLeft() {
		n.Left.ForEach(operation)
	}
	if n.hasRight() {
		n.Right.ForEach(operation)
	}
}
