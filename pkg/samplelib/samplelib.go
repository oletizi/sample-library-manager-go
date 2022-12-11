package samplelib

type Sample struct {
	name string
	path string
	null bool
}

func NewSample(name string, path string) Sample {
	return Sample{name: name, path: path}
}

type Node struct {
	name     string
	path     string
	parent   *Node
	children []*Node
	samples  []*Sample
	null     bool
}

func NullNode() *Node {
	rv := &Node{name: "", path: "", children: []*Node{}, samples: []*Sample{}, null: true}
	rv.parent = rv
	return rv
}

func NewNode(name string, path string, parent *Node, children []*Node, samples []*Sample) *Node {
	return &Node{name: name, path: path, parent: parent, children: children, samples: samples, null: false}
}
