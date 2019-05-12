package main

const (
	asciiCharsCount = 256
)

var (
	asciiSpaces = [asciiCharsCount]bool{'\t': true, '\n': true, '\v': true, '\f': true, '\r': true, ' ': true}
)

type Word struct {
	Value []byte
	Count int
}

type Dictionary []Word

func (d Dictionary) Len() int           { return len(d) }
func (d Dictionary) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d Dictionary) Less(i, j int) bool { return d[i].Count < d[j].Count }

func NewNode(value byte) *Node {
	return &Node{Value: value, Children: [asciiCharsCount]*Node{}}
}

type Node struct {
	Value    byte
	Count    int
	Children [asciiCharsCount]*Node
	IsLast   bool
}

func (n *Node) Store(value byte) *Node {
	node := n.Children[value]
	if node != nil {
		return node
	}

	node = NewNode(value)
	n.Children[value] = node

	return node
}

func NewPrefixTree() *PrefixTree {
	return &PrefixTree{Root: NewNode(0)}
}

type PrefixTree struct {
	Root *Node
}

func (p *PrefixTree) Insert(word []byte) {
	node := p.Root

	for _, b := range word {
		node = node.Store(b)
	}

	node.IsLast = true
	node.Count++
	return
}

func (p *PrefixTree) ListAll() Dictionary {
	dict := &Dictionary{}
	p.traverse([]byte{}, p.Root, dict)

	return *dict
}

func (p *PrefixTree) traverse(prefix []byte, node *Node, dict *Dictionary) {
	if node.IsLast {
		value := make([]byte, len(prefix))
		copy(value, prefix)
		*dict = append(*dict, Word{Value: value, Count: node.Count})
	}

	for _, child := range node.Children {
		if child == nil {
			continue
		}

		prefix = append(prefix, child.Value)
		p.traverse(prefix, child, dict)
		prefix = prefix[:len(prefix)-1]
	}
}
