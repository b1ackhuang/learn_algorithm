package rb_tree

const (
	Red   Color = "r"
	Black Color = "b"
)

type Color string

type RBNode struct {
	Color
	val    int
	Parent *RBNode
	Left   *RBNode
	Right  *RBNode
}

func NewNode(val int) *RBNode {
	return &RBNode{Color: Red}
}

func (r *RBNode) isRoot() bool {
	return r.Parent == nil
}

func (r *RBNode) isLeftSon() bool {
	return !r.isRoot() && r.Parent.Left == r
}

func (r *RBNode) isRightSon() bool {
	return !r.isRoot() && r.Parent.Right == r
}
