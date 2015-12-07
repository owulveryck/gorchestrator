package structure

import (
	"encoding/json"
	"fmt"
)

func ExampleInput() {
	test := Input{"Test",
		[]int{0, 1, 0, 0, 1, 0, 0, 0,
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
	o, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", o)
}
