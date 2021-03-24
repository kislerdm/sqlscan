package main

import (
	_ "embed"
	"log"

	"github.com/kislerdm/sqlscan/graph"
	"github.com/kislerdm/sqlscan/graph/viz"
	"github.com/kislerdm/sqlscan/internal/fs"
)

func main() {
	g := graph.Graph{
		Nodes: []graph.Node{
			{ID: "schema1.table1", Name: "table1", Group: "schema1", Type: "table"},
			{ID: "schema2.table2", Name: "table2", Group: "schema2", Type: "table"},
		},
		Edges: []graph.Edge{
			{ID: "1", Source: "schema1.table1", Target: "schema2.table2"},
		},
	}
	_, err := viz.ToHTML(g)
	if err != nil {
		log.Fatalln(err)
	}

	err = fs.WriteFile([]byte{}, "/tmp/index.html")
	if err != nil {
		log.Fatalln(err)
	}
}
