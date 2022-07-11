package tire_tree

type Trie struct {
	root TrieNode
}

type TrieNode struct {
	val      byte
	sonNodes []TrieNode
}

func Constructor() Trie {
	return Trie{root: TrieNode{}}
}

func (this *Trie) Insert(word string) {
	var parentNode = this.root
	for i := 0; i < len(word); i++ {
		var node = TrieNode{val: word[i]}
		var curNode, curSonNodes = insertNodes(node, parentNode.sonNodes)
		parentNode.sonNodes = curSonNodes
		parentNode = *curNode
	}
}

func insertNodes(node TrieNode, l []TrieNode) (*TrieNode, []TrieNode) {
	var insertIndex = 0
	for insertIndex < len(l) && node.val <= l[insertIndex].val {
		if node.val == l[insertIndex].val {
			return &l[insertIndex], l
		}
		insertIndex++
	}
	var retList = make([]TrieNode, 0, len(l)+1)
	retList = append(retList, l[:insertIndex]...)
	retList = append(retList, node)
	retList = append(retList, l[:insertIndex]...)
	return &l[insertIndex], retList
}

func existsNode(val byte, l []TrieNode) *TrieNode {
	for i := 0; i < len(l); i++ {
		if l[i].val == val {
			return &l[i]
		}
	}
	return nil
}

func (this *Trie) Search(word string) bool {
	var parentNodes = &this.root
	for i := 0; i < len(word); i++ {
		var tmpCurNode = existsNode(word[i], parentNodes.sonNodes)
		if tmpCurNode == nil {
			return false
		}
		parentNodes = tmpCurNode
	}
	return parentNodes != nil
}

func (this *Trie) StartsWith(prefix string) bool {
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
