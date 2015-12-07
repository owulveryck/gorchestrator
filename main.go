package main

import (
	"github.com/owulveryck/gorchestrator/orchestrator"
)

var valid orchestrator.Input
var notValid orchestrator.Input

func init() {
	valid = orchestrator.Input{"Valid",
		[]int64{0, 1, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 1, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
		},
		[]orchestrator.Node{
			{0, "a", "shell", "example/script.sh", nil, nil},
			{1, "b", "shell", "myscript.sh", nil,
				map[string]string{
					"output1": "",
				},
			},
			{2, "c", "shell", "example/script.sh",
				[]string{
					"-e", "get_attribute 1:output1",
				}, nil},
			{3, "d", "shell", "example/script.sh", nil, nil},
			{4, "e", "shell", "example/script.sh", nil, nil},
			{5, "f", "shell", "example/script.sh", nil, nil},
			{6, "g", "shell", "example/script.sh", nil, nil},
			{7, "h", "shell", "example/script.sh", nil, nil},
		},
	}

	notValid = orchestrator.Input{"NotValid",
		[]int64{0, 1, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 1, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
		},
		[]orchestrator.Node{
			{0, "a", "shell", "example/script.sh", nil, nil},
			{1, "b", "shell", "myscript.sh", nil,
				map[string]string{
					"output1": "",
				},
			},
			{2, "c", "shell", "myscript2.sh",
				[]string{
					"-e", "get_attribute 1:output1",
				}, nil},
			{3, "d", "shell", "example/script.sh", nil, nil},
			{4, "e", "shell", "example/script.sh", nil, nil},
			{5, "f", "shell", "example/script.sh", nil, nil},
			{6, "g", "shell", "example/script.sh", nil, nil},
			{7, "h", "shell", "example/script.sh", nil, nil},
		},
	}
}
func main() {
	valid.Run()
}
