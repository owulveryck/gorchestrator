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
	//"encoding/json"
	"flag"
	//"fmt"
	"github.com/owulveryck/gorchestrator/orchestrator"
	"github.com/owulveryck/toscalib"
	//"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var t toscalib.ServiceTemplateDefinition
	var v orchestrator.Graph
	flag.Parse()
	if toscaFilename == "" {
		flag.PrintDefaults()
		return
	}
	if inputFilename == "" {
		log.Println("Warning: No input file passed as argument, using default values")
	}

	inputs, err := getInputs(inputFilename)
	log.Println(inputs)

	r, err := os.Open(toscaFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	// Change the CWD to deal correctly with the imports of the TOSCA file
	os.Chdir(filepath.Dir(toscaFilename))
	err = t.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	v = togorch(t)
	for _, n := range v.Nodes {
		log.Printf("[%v]\t%v\t%v\t%v", n.Name, n.Artifact, n.Args, n.Outputs)

	}
	//res, _ := json.MarshalIndent(v, "  ", "  ")
	//fmt.Printf("%s\n", string(res))
}
