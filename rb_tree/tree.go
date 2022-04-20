package rb_tree

type RBTree struct {
	Root *RBNode
}

func (r *RBTree) Insert(val int) {
	var node = NewNode(val)
	if r.Root == nil {
		node.Color = Black
		r.Root = node
		return
	}

	//1 查找待插入节点
	var cur = r.Root
	for cur.Left != nil && cur.Right != nil {
		if cur.val == val {
			return
		} else if val < cur.Left.val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	//case1：父节点的颜色为黑色
	if cur.Color == Black {
		if val < cur.val {
			cur.Left = node
		} else {
			cur.Right = node
		}
		return
	}

	//case2：父节点的颜色为红色

}

func rotateLeft(r *RBNode) {
	var parent, son = r.Parent, r.Right
	if r.isLeftSon() {
		parent.Left = son
	} else {
		parent.Right = son
	}

	r.Right = son.Left

	son.Left = r
	r.Parent = son

}

func rotateRight(r *RBNode) {
	var parent, son = r.Parent, r.Right
	if r.isLeftSon() {
		parent.Left = son
	} else {
		parent.Right = son
	}

	r.Left = son.Right

	son.Right = r
	r.Parent = son
}
