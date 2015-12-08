package orchestrator

import (
	"encoding/json"
	"fmt"
	"testing"
)

var valid Graph
var notValid Graph

func init() {
	valid = Graph{"Valid",
		0,
		[]int64{0, 1, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 1, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
		},
		[]Node{
			{0, "a", "nil", "myplaybook.yml", nil, nil},
			{1, "b", "shell", "myscript.sh", nil,
				map[string]string{
					"output1": "",
				},
			},
			{2, "c", "shell", "myscript2.sh",
				[]string{
					"-e", "get_attribute 1:output1",
				}, nil},
			{3, "d", "nil", "myplaybook3.yml", nil, nil},
			{4, "e", "nil", "myplaybook4.yml", nil, nil},
			{5, "f", "nil", "myplaybook5.yml", nil, nil},
			{6, "g", "nil", "myplaybook6.yml", nil, nil},
			{7, "h", "nil", "myplaybook7.yml", nil, nil},
		},
	}

	notValid = Graph{"NotValid",
		0,
		[]int64{0, 1, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 1, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
		},
		[]Node{
			{0, "a", "nil", "myplaybook.yml", nil, nil},
			{1, "b", "shell", "myscript.sh", nil,
				map[string]string{
					"output1": "",
				},
			},
			{2, "c", "shell", "myscript2.sh",
				[]string{
					"-e", "get_attribute 1:output1",
				}, nil},
			{3, "d", "nil", "myplaybook3.yml", nil, nil},
			{4, "e", "nil", "myplaybook4.yml", nil, nil},
			{5, "f", "nil", "myplaybook5.yml", nil, nil},
			{6, "g", "nil", "myplaybook6.yml", nil, nil},
			{7, "h", "nil", "myplaybook7.yml", nil, nil},
		},
	}
}

func TestCheck(t *testing.T) {
	e := valid.Check()
	if e.Code != 0 {
		t.Errorf("Struct should be valid, error is: %v", e.Error())
	}
	e = notValid.Check()
	if e.Code == 0 {
		t.Errorf("Struct should not be valid, error is: %v", e.Error())
	}
}

func ExampleCheck() {
	test := Graph{"Test",
		0,
		[]int64{0, 1, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 1, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
		},
		[]Node{
			{0, "a", "nil", "myplaybook.yml", nil, nil},
			{1, "b", "shell", "myscript.sh", nil,
				map[string]string{
					"output1": "",
				},
			},
			{2, "c", "shell", "myscript2.sh",
				[]string{
					"-e", "get_attribute 1:output1",
				}, nil},
			{3, "d", "nil", "myplaybook3.yml", nil, nil},
			{4, "e", "nil", "myplaybook4.yml", nil, nil},
			{5, "f", "nil", "myplaybook5.yml", nil, nil},
			{6, "g", "nil", "myplaybook6.yml", nil, nil},
			{7, "h", "nil", "myplaybook7.yml", nil, nil},
		},
	}
	e := test.Check()
	if e.Code != 0 {
		panic(e.Error)
	}
	o, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", o)
}
