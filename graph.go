package main

// START_INT OMIT
type NodeInterface interface {
	Edges() []EdgeInterface
}

type EdgeInterface interface {
	Nodes() (from, to NodeInterface)
}

type Graph struct { ... }

func New(nodes []NodeInterface) *Graph {
	...
}

func (g *Graph) ShortestPath(from, to NodeInterface) []EdgeInterface { ... }
// END_INT OMIT

// START_INT_IMPL OMIT
type Vertex struct { ... }

func (v *Vertex) Edges() []EdgeInterface { ... }

type FromTo struct { ... }

func (e *FromTo) Nodes() (from, to NodeInterface) { ... }
// END_INT_IMPL OMIT

// START_INT_EXAMPLE OMIT
vertices := []*Vertex{...}
nodes := make([]NodeInterface, 0, len(vertices))
for _, v := range vertices {
	nodes = append(nodes, v)
}

graph := New(nodes)
var edges []EdgeInterface = graph.ShortestPath(from, to)

path := make([]*FromTo, 0, len(edges))
for _, e := range edges {
	path = append(path, e.(*FromTo))
}
// END_INT_EXAMPLE OMIT

// START_TP OMIT
type NodeConstraint[Edge any] interface {
	Edges() []Edge
}

type EdgeConstraint[Node any] interface {
	Nodes() (from, to Node)
}

type Graph[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]] struct { ... }

func New[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]](nodes []Node) *Graph[Node, Edge] {
	...
}

func (g *Graph[Node, Edge]) ShortestPath(from, to Node) []Edge{ ... }
// END_TP OMIT

// START_TP_IMPL OMIT
type Vertex struct { ... }

func (v *Vertex) Edges() []*FromTo { ... }

type FromTo struct { ... }

func (e *FromTo) Nodes() (from, to *Vertex) { ... }

// END_TP_IMPL OMIT

// START_TP_EXAMPLE OMIT
vertices := []*Vertex{...}

graph := New[*Vertex, *FromTo](vertices)
var path []*FromTo = graph.ShortestPath(from, to)
// END_TP_EXAMPLE OMIT
