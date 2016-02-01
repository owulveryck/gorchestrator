/*
Olivier Wulveryck - author of Gorchestrator
Copyright (C) 2015 Olivier Wulveryck

This file is part of the Gorchestrator project and
is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
	"strings"
	"text/template"
)

var svg map[string][]byte

func getGraph(id string) (orchestrator.Graph, error) {
	var g orchestrator.Graph

	r, err := http.Get(fmt.Sprintf("%v/%v", OrchestratorUrl, id))
	if err != nil {
		return g, err
	}
	defer r.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&g); err != nil {
		return g, err
	}

	return g, nil
}

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

	template, err := template.New("node").Parse(`{{define "NODE"}}<<table border="0" cellspacing="0">
		<tr ><td colspan="2" port="port1" border="1" bgcolor="lightblue">{{.Name}}</td></tr>
		<tr ><td colspan="2" port="port2" border="1">{{.Target}}</td></tr>
		<tr>
			<td port="port2" border="1">{{.Engine}}</td>
			<td port="port8" border="1">{{.Artifact}}</td>
		</tr>
		{{range .Args}}
		<tr ><td colspan="2" port="port2" border="1">{{.}}</td></tr>
		{{end}}
		<tr ><td colspan="2" port="port2" border="1">{{.Outputs}}</td></tr>
		</table>>{{end}}`)

	// Creates a new graph
	g := gographviz.NewGraph()
	//g.AddAttr("", "rankdir", "LR")
	// Now read the json input
	var v orchestrator.Graph

	resp, err := http.Get(fmt.Sprintf("%v/%v", OrchestratorUrl, id))
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
	g.SetName("Execution")
	g.SetDir(true)
	m := make(map[int]string)
	// Now add every node
	// Hack: If node has no Engine and no Artiface, and if sum(row)=sum(col)=0, skip it
	for i, _ := range v.Nodes {
		sumr := int64(0)
		sumc := int64(0)
		for r := 0; r < v.Digraph.Dim(); r++ {
			sumr = sumr + v.Digraph.At(v.Nodes[i].ID, r)
			sumc = sumc + v.Digraph.At(r, v.Nodes[i].ID)
		}
		if v.Nodes[i].Artifact == "" && sumr == 0 && sumc == 0 {
			continue
		}
		tmp := make([]string, 2)
		v.Nodes[i].Name = strings.Replace(v.Nodes[i].Name, "-", "_", -1)
		tmp = strings.SplitAfter(v.Nodes[i].Name, ":")
		v.Nodes[i].Name = strings.Replace(v.Nodes[i].Name, ":", "_", -1)
		if len(tmp) != 2 {
			tmp[0] = v.Nodes[i].Name
			tmp = append(tmp, "")
		}
		var out bytes.Buffer
		err = template.ExecuteTemplate(&out, "NODE", v.Nodes[i])
		g.AddNode("G", v.Nodes[i].Name,
			map[string]string{
				"id": fmt.Sprintf("\"%v\"", strconv.Itoa(v.Nodes[i].ID)),
				//"label": fmt.Sprintf("\"%v|%v\"", tmp[0], tmp[1]),
				"label": out.String(),
				//"label": fmt.Sprintf("\"%v|%v|%v|%v|%v\"", tmp[0], tmp[1], v.Nodes[i].Engine, v.Nodes[i].Artifact, v.Nodes[i].Args[:]),
				"shape": "\"record\"",
				"style": "\"rounded\"",
			})
		m[v.Nodes[i].ID] = v.Nodes[i].Name
	}
	for r := 0; r < v.Digraph.Dim(); r++ {
		for c := 0; c < v.Digraph.Dim(); c++ {
			if v.Digraph.At(r, c) != 0 {
				//g.AddEdge(fmt.Sprintf("%v:%v", m[r], r), fmt.Sprintf("%v:%v", m[c], c), true, nil)
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

	err = d.Wait()
	if err != nil {
		fmt.Println(s)
	}

	return buf.Bytes(), nil
}
