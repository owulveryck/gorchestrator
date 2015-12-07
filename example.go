package main

import (
	"github.com/owulveryck/gorchestrator/structure"
)

var valid structure.Input
var notValid structure.Input

func init() {
	valid = structure.Input{"Valid",
		[]int64{0, 1, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 1, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
		},
		[]structure.Node{
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

	notValid = structure.Input{"NotValid",
		[]int64{0, 1, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 1, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
		},
		[]structure.Node{
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
