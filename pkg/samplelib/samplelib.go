package samplelib

type Sample struct {
	Name string
	Path string
	null bool
}

func NewSample(name string, path string) *Sample {
	return &Sample{Name: name, Path: path}
}

type Node struct {
	Name   string
	Path   string
	Parent *Node
	null   bool
}

func NullNode() *Node {
	rv := &Node{Name: "", Path: "", null: true}
	rv.Parent = rv
	return rv
}

type DataSource interface {
	RootNode() (*Node, error)
	ChildrenOf(node *Node) ([]*Node, error)
	SamplesOf(node *Node) ([]*Sample, error)
}
