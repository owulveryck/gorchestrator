package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/awalterschulze/gographviz"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

var svg map[string][]byte

func getSvg(id string) ([]byte, error) {
	if b, ok := svg[id]; ok {
		return b, nil
	} else {
		b, err := generateSvg(id)
		if err != nil {
			return nil, err
		} else {
			return b, nil
		}
	}
}

func generateSvg(id string) ([]byte, error) {
	// Creates a new graph
	g := gographviz.NewGraph()
	g.AddAttr("", "rankdir", "LR")
	// Now read the json input
	var v orchestrator.Graph

	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/v1/tasks/%v", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}

	// Now for each node, create a node
	g.SetName(v.Name)
	g.SetDir(true)
	m := make(map[int]string)
	// Now add every node
	for _, n := range v.Nodes {
		g.AddNode("G", n.Name,
			map[string]string{
				"id":    fmt.Sprintf("\"%v\"", strconv.Itoa(n.ID)),
				"label": fmt.Sprintf("\"%v|%v|%v\"", n.Name, n.Engine, n.Artifact),
				"shape": "\"record\"",
			})
		m[n.ID] = n.Name
	}
	for r := 0; r < v.Digraph.Dim(); r++ {
		for c := 0; c < v.Digraph.Dim(); c++ {
			if v.Digraph.At(r, c) != 0 {
				g.AddEdge(m[r], m[c], true, nil)
			}
		}
	}
	// Now add the edges
	s := g.String()
	d := exec.Command("dot", "-Tsvg")

	// Set the stdin stdout and stderr of the dot subprocess
	stdinOfDotProcess, err := d.StdinPipe()
	if err != nil {
		return nil, err

	}
	defer stdinOfDotProcess.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line
	readCloser, err := d.StdoutPipe()
	if err != nil {
		return nil, err

	}
	d.Stderr = os.Stderr

	// Actually run the dot subprocess
	if err = d.Start(); err != nil { //Use start, not run
		fmt.Println("An error occured: ", err) //replace with logger, or anything you want
	}
	fmt.Fprintf(stdinOfDotProcess, s)
	stdinOfDotProcess.Close()

	// Read from stdout and store it in the correct structure
	var buf bytes.Buffer
	buf.ReadFrom(readCloser)

	d.Wait()
	return buf.Bytes(), nil
}
