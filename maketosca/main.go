package main

import (
	"flag"
	"github.com/owulveryck/toscalib"
	"github.com/owulveryck/gorschestrator"
	"github.com/owulveryck/toscaviewer"
	"log"
	"net/http"
	"sync"
)

func main() {


	// Fet the rooted path name of the current directory
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	example := fmt.Sprintf("%v/../examples/tosca_single_instance_wordpress.yaml", pwd)
	example = filepath.Clean(example)
	var testFile = flag.String("testfile", example, "a tosca yaml file to process")
	flag.Parse()

	var toscaTemplate toscalib.ToscaDefinition
	file, err := os.Open(*testFile)

	if err != nil {
		log.Panic("error: ", err)
	}
	//err = yaml.Unmarshal(file, &toscaTemplate)
	err = toscaTemplate.Parse(file)
	if err != nil {
		log.Panic("error: ", err)
	}
	router := toscaviewer.NewRouter(&toscaTemplate)

	log.Println("connect here: http://localhost:8080/svg")
	log.Fatal(http.ListenAndServe(":8080", router))
	//taskStructure.PrintAdjacencyMatrix()
	// Entering the workers area
	/*
	doneChan := make(chan *gorchestrator.Task)

	// For each task, launch a goroutine
	for taskIndex, _ := range taskStructure.Tasks {
		go gorchestrator.Runner(taskStructure.Tasks[taskIndex], doneChan, &wg)
		wg.Add(1)
	}
	go gorchestrator.Advertize(taskStructure, doneChan)

	// This is the web displa
	router := gorchestrator.NewRouter(taskStructure)

	go log.Fatal(http.ListenAndServe(":8080", router))

	// Wait for all the runner(s) to be finished
	wg.Wait()
	*/
}
