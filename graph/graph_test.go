package graph_test

import (
	"reflect"
	"testing"

	"github.com/kislerdm/sqlscan/graph"
)

func TestToJSON(t *testing.T) {
	tests := []struct {
		in   *graph.Graph
		want []byte
	}{
		{
			in: &graph.Graph{
				Nodes: []graph.Node{
					{ID: "schema1.table1", Name: "table1", Group: "schema1", Type: "table"},
					{ID: "schema2.table2", Name: "table2", Group: "schema2", Type: "table"},
				},
				Edges: []graph.Edge{
					{ID: "1", Source: "schema1.table1", Target: "schema2.table2"},
				},
			},
			want: []byte(`{"nodes":[{"id":"schema1.table1","name":"table1","group":"schema1","type":"table"},{"id":"schema2.table2","name":"table2","group":"schema2","type":"table"}],"edges":[{"id":"1","source":"schema1.table1","target":"schema2.table2"}]}`),
		},
	}
	for _, test := range tests {
		got, _ := test.in.ToJSON()
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("Results don't match\nwant: %v\ngot: %v,\n%s", test.want, got, string(got))
		}
	}
}

func TestFromJSON(t *testing.T) {
	tests := []struct {
		in   []byte
		want graph.Graph
	}{
		{
			in: []byte(`{"nodes":[{"id":"schema1.table1","name":"table1","group":"schema1","type":"table"},{"id":"schema2.table2","name":"table2","group":"schema2","type":"table"}],"edges":[{"id":"1","source":"schema1.table1","target":"schema2.table2"}]}`),
			want: graph.Graph{
				Nodes: []graph.Node{
					{ID: "schema1.table1", Name: "table1", Group: "schema1", Type: "table"},
					{ID: "schema2.table2", Name: "table2", Group: "schema2", Type: "table"},
				},
				Edges: []graph.Edge{
					{ID: "1", Source: "schema1.table1", Target: "schema2.table2"},
				},
			},
		},
	}
	for _, test := range tests {
		got, _ := graph.FromJSON(test.in)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("Results don't match\nwant: %v\ngot: %v", test.want, got)
		}
	}
}
