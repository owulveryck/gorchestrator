package orchestrator

import (
	"encoding/json"
	"fmt"
	"testing"
)

var valid Input
var notValid Input

func init() {
	valid = Input{"Valid",
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
			{0, "a", "ansible", "myplaybook.yml", nil, nil},
			{1, "b", "shell", "myscript.sh", nil,
				map[string]string{
					"output1": "",
				},
			},
			{2, "c", "shell", "myscript2.sh",
				[]string{
					"-e", "get_attribute 1:output1",
				}, nil},
			{3, "d", "ansible", "myplaybook3.yml", nil, nil},
			{4, "e", "ansible", "myplaybook4.yml", nil, nil},
			{5, "f", "ansible", "myplaybook5.yml", nil, nil},
			{6, "g", "ansible", "myplaybook6.yml", nil, nil},
			{7, "h", "ansible", "myplaybook7.yml", nil, nil},
		},
	}

	notValid = Input{"NotValid",
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
			{0, "a", "ansible", "myplaybook.yml", nil, nil},
			{1, "b", "shell", "myscript.sh", nil,
				map[string]string{
					"output1": "",
				},
			},
			{2, "c", "shell", "myscript2.sh",
				[]string{
					"-e", "get_attribute 1:output1",
				}, nil},
			{3, "d", "ansible", "myplaybook3.yml", nil, nil},
			{4, "e", "ansible", "myplaybook4.yml", nil, nil},
			{5, "f", "ansible", "myplaybook5.yml", nil, nil},
			{6, "g", "ansible", "myplaybook6.yml", nil, nil},
			{7, "h", "ansible", "myplaybook7.yml", nil, nil},
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
	test := Input{"Test",
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
			{0, "a", "ansible", "myplaybook.yml", nil, nil},
			{1, "b", "shell", "myscript.sh", nil,
				map[string]string{
					"output1": "",
				},
			},
			{2, "c", "shell", "myscript2.sh",
				[]string{
					"-e", "get_attribute 1:output1",
				}, nil},
			{3, "d", "ansible", "myplaybook3.yml", nil, nil},
			{4, "e", "ansible", "myplaybook4.yml", nil, nil},
			{5, "f", "ansible", "myplaybook5.yml", nil, nil},
			{6, "g", "ansible", "myplaybook6.yml", nil, nil},
			{7, "h", "ansible", "myplaybook7.yml", nil, nil},
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
