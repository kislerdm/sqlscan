package graph

import "encoding/json"

// Node defines a graph node.
type Node struct {
	// Node id, i.e. {{schema}}.{{name}}
	ID string `json:"id"`
	// Node name, i.e. table name
	Name string `json:"name"`
	// Node group, i.e. schema
	Group string `json:"group"`
	// Node type, e.g. table, or view
	Type string `json:"type"`
}

// Edge defines a graph edge.
type Edge struct {
	ID string `json:"id"`
	// ID on the incoming node
	Source string `json:"source,omitempty"`
	// ID on the outgoing node
	Target string `json:"target,omitempty"`
}

// Graph defines a graph.
type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

// ToJSON serializes the graph object.
func (g *Graph) ToJSON() ([]byte, error) {
	return json.Marshal(g)
}

// FromJSON de-serializes the graph object.
func FromJSON(data []byte) (Graph, error) {
	var g Graph
	err := json.Unmarshal(data, &g)
	return g, err
}
