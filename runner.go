package gorschestrator

import (
	//"github.com/gonum/matrix/mat64" // Matrix
	"log"
	"math/rand" // Temp
	"github.com/owulveryck/toscalib"
	"sync"
	//"time"
)

func random(min int, max int) int {
	var bytes int
	bytes = min + rand.Intn(max)
	return int(bytes)
	//rand.Seed(time.Now().UTC().UnixNano())
	//return rand.Intn(max - min) + min
}

// Runner runs the lifecycle of node
// it waits for a boolean in the channel to see if it may run
func Runner(node *toscalib.NodeTemplate, doneChan chan<- *toscalib.NodeTemplate, wg *sync.WaitGroup) {
	//log.Printf("[%v:%v] Queued", task.Id, task.Name)
	for {
		whatToDo := <-node.RunChan
		// For each dependency of the task
		// We can run if the sum of the element of the column Id of the current task is 0

		switch whatToDo {
		case toscalib.StateConfiguring:
			log.Println("Configuring")
		}
		/*
		if letsGo == toscalib.StateConfiguring {
			proto := "tcp"
			socket := task.Node
			task.Status = -1
			log.Printf("[%v:%v] Running (%v)", task.Id, task.Name, task.Module)
			log.Printf("[%v] Connecting in %v on %v", task.Name, proto, socket)
			task.StartTime = time.Now()
			if task.Module != "dummy" && task.Module != "meta" && task.Node != "null" {
				log.Printf("Sending command on %v", task.Node)
				task.Status = Client(task, &proto, &socket)
			} else {
				task.Status = 0
			}
			task.EndTime = time.Now()
			// ... Do a lot of stufs...
			//time.Sleep(time.Duration(sleepTime) * time.Second)
			// Adjust the Status
			//task.Status = 2
			// Send it on the channel
			log.Printf("[%v:%v] Done", task.Id, task.Name)
			doneChan <- task
			wg.Done()
			return
		}
		*/
	}
}
/*
// Advertize goroutine, reads the tasks from doneChannel and write the TaskGraphStructure back to the taskStructureChan
func Advertize(taskStructure *TaskGraphStructure, doneChan <-chan *Task) {
	// Let's launch the task that can initially run
	rowSize, _ := taskStructure.AdjacencyMatrix.Dims()
	for taskIndex, _ := range taskStructure.Tasks {
		sum := float64(0)
		for r := 0; r < rowSize; r++ {
			sum += taskStructure.AdjacencyMatrix.At(r, taskIndex)
		}
		if sum == 0 && taskStructure.Tasks[taskIndex].Status < 0 {
			taskStructure.Tasks[taskIndex].TaskCanRunChan <- true
		}
	}
	doneAdjacency := mat64.DenseCopyOf(taskStructure.AdjacencyMatrix)
	// Store the task that we have already advertized
	var advertized []int
	for {
		task := <-doneChan

		// TaskId is finished, it cannot be the source of any task anymore
		// Set the row at 0 if status is 0
		rowSize, colSize := doneAdjacency.Dims()
		if task.Status == 0 {
			for c := 0; c < colSize; c++ {
				doneAdjacency.Set(task.Id, c, float64(0))
			}
		}
		// For each dependency of the task
		// We can run if the sum of the element of the column Id of the current task is 0
		for taskIndex, _ := range taskStructure.Tasks {
			sum := float64(0)
			for r := 0; r < rowSize; r++ {
				sum += doneAdjacency.At(r, taskIndex)
			}

			// This task can be advertized...
			if sum == 0 && taskStructure.Tasks[taskIndex].Status < -2 {
				taskStructure.Tasks[taskIndex].Status = -2
				// ... if it has not been advertized already
				advertized = append(advertized, taskIndex)
				taskStructure.Tasks[taskIndex].TaskCanRunChan <- true
			}
		}
	}
}
*/
