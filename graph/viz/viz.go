package viz

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"html/template"

	"github.com/kislerdm/sqlscan/graph"
)

// elements defines the graph elements for cytoscape.js.
// See https://js.cytoscape.org/ for details.
type elements []data

type data struct {
	Data map[string]string `json:"data"`
}

func fromGraphElement(e interface{}) ([]data, error) {
	var o []map[string]string
	if n, err := json.Marshal(e); err != nil {
		return []data{}, err
	} else {
		if err := json.Unmarshal(n, &o); err != nil {
			return []data{}, err
		}
	}
	var output []data
	for _, el := range o {
		output = append(output, data{Data: el})
	}
	return output, nil
}

func fromGraph(g graph.Graph) (elements, error) {
	nodes, err := fromGraphElement(g.Nodes)
	if err != nil {
		return elements{}, err
	}
	edges, err := fromGraphElement(g.Edges)
	if err != nil {
		return elements{}, err
	}
	return append(nodes, edges...), nil
}

var (
	//go:embed template.html
	htmlTemplate string
	//go:embed cytoscape.min.js
	lib string
	//go:embed klay.js
	layout string
)

type Template struct {
	Elements elements
	Lib      string
	Layout   string
}

func ToHTML(g graph.Graph) (string, error) {
	elements, err := fromGraph(g)
	if err != nil {
		return "", err
	}

	t := Template{
		Elements: elements,
		Lib:      lib,
		Layout:   layout,
	}

	p, err := template.New("template").Parse(htmlTemplate)
	if err != nil {
		return "", err
	}

	var templateRender bytes.Buffer
	if err := p.Execute(&templateRender, t); err != nil {
		return "", err
	}
	return templateRender.String(), nil
}
