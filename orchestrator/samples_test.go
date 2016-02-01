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

package orchestrator

import (
	"github.com/owulveryck/gorchestrator/structure"
	"time"
)

var valid Graph
var validAndTimeout Graph
var validAndNoArtifact Graph
var validAndFailed Graph
var validAndSleep Graph
var validAndSleepAndFailed Graph
var validAndExecSuccess Graph
var validAndExecFailure Graph
var notValid Graph

func init() {
	matrix := structure.Matrix{
		Matrix: []int64{
			0, 1, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 1, 0, 0, 0, 1, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
		},
	}

	validAndNoArtifact = Graph{"ValidAndNoArtifact",
		0,
		matrix,
		[]Node{
			{0, 0, "a", "SELF", "nil", "", nil, nil, "", "", nil},
			{1, 0, "b", "SELF", "nil", "", nil, nil, "", "", nil},
			{2, 0, "c", "SELF", "nil", "", nil, nil, "", "", nil},
			{3, 0, "d", "SELF", "nil", "", nil, nil, "", "", nil},
			{4, 0, "e", "SELF", "nil", "", nil, nil, "", "", nil},
			{5, 0, "f", "SELF", "nil", "", nil, nil, "", "", nil},
			{6, 0, "g", "SELF", "nil", "", nil, nil, "", "", nil},
			{7, 0, "h", "SELF", "nil", "", nil, nil, "", "", nil},
		},
		time.After(30 * time.Second),
		"",
	}
	validAndExecSuccess = Graph{"ValidAndExecSuccess",
		0,
		matrix,
		[]Node{
			{0, 0, "a", "", "shell", "success", nil, nil, "", "", nil},
			{1, 0, "b", "", "", "nil", nil, nil, "", "", nil},
			{2, 0, "c", "", "", "", nil, nil, "", "", nil},
			{3, 0, "d", "", "shell", "success", nil, nil, "", "", nil},
			{4, 0, "e", "", "", "", nil, nil, "", "", nil},
			{5, 0, "f", "", "", "", nil, nil, "", "", nil},
			{6, 0, "g", "", "", "", nil, nil, "", "", nil},
			{7, 0, "h", "", "", "", nil, nil, "", "", nil},
		},
		time.After(30 * time.Second),
		"",
	}
	validAndExecFailure = Graph{"ValidAndExecFailure",
		0,
		matrix,
		[]Node{
			{0, 0, "a", "", "shell", "success", nil, nil, "", "", nil},
			{1, 0, "b", "", "", "nil", nil, nil, "", "", nil},
			{2, 0, "c", "", "", "", nil, nil, "", "", nil},
			{3, 0, "d", "", "shell", "failure", nil, nil, "", "", nil},
			{4, 0, "e", "", "", "", nil, nil, "", "", nil},
			{5, 0, "f", "", "", "", nil, nil, "", "", nil},
			{6, 0, "g", "", "", "", nil, nil, "", "", nil},
			{7, 0, "h", "", "", "", nil, nil, "", "", nil},
		},
		time.After(30 * time.Second),
		"",
	}

	validAndTimeout = Graph{"ValidAndTimeout",
		0,
		matrix,
		[]Node{
			{0, 0, "a", "", "sleep", "", nil, nil, "", "", nil},
			{1, 0, "b", "", "sleep", "", nil, nil, "", "", nil},
			{2, 0, "c", "", "sleep", "", nil, nil, "", "", nil},
			{3, 0, "d", "", "sleep", "", nil, nil, "", "", nil},
			{4, 0, "e", "", "sleep", "", nil, nil, "", "", nil},
			{5, 0, "f", "", "sleep", "", nil, nil, "", "", nil},
			{6, 0, "g", "", "sleep", "", nil, nil, "", "", nil},
			{7, 0, "h", "", "sleep", "", nil, nil, "", "", nil},
		},
		time.After(1 * time.Millisecond),
		"",
	}

	validAndSleep = Graph{"ValidAndSleep",
		0,
		matrix,
		[]Node{
			{0, 0, "a", "", "sleep", "", nil, nil, "", "", nil},
			{1, 0, "b", "", "sleep", "", nil, nil, "", "", nil},
			{2, 0, "c", "", "sleep", "", nil, nil, "", "", nil},
			{3, 0, "d", "", "sleep", "", nil, nil, "", "", nil},
			{4, 0, "e", "", "sleep", "", nil, nil, "", "", nil},
			{5, 0, "f", "", "sleep", "", nil, nil, "", "", nil},
			{6, 0, "g", "", "sleep", "", nil, nil, "", "", nil},
			{7, 0, "h", "", "sleep", "", nil, nil, "", "", nil},
		},
		time.After(30 * time.Second),
		"",
	}

	valid = Graph{"Valid",
		0,
		matrix,
		[]Node{
			{0, 0, "a", "", "nil", "myplaybook.yml", nil, nil, "", "", nil},
			{1, 0, "b", "", "nil", "myscript.sh", nil,
				map[string]string{
					"result": "",
				},
				"", "", nil},
			{2, 0, "c", "", "nil", "myscript2.sh",
				[]string{
					"-e", "get_attribute b:result",
				}, nil, "", "", nil},
			{3, 0, "d", "", "nil", "myplaybook3.yml", nil, nil, "", "", nil},
			{4, 0, "e", "", "nil", "myplaybook4.yml", nil, nil, "", "", nil},
			{5, 0, "f", "", "nil", "myplaybook5.yml", nil, nil, "", "", nil},
			{6, 0, "g", "", "nil", "myplaybook6.yml", nil, nil, "", "", nil},
			{7, 0, "h", "", "nil", "myplaybook7.yml", nil, nil, "", "", nil},
		},
		time.After(30 * time.Second),
		"",
	}

	notValid = Graph{"NotValid",
		0,
		structure.Matrix{
			Matrix: []int64{0, 1, 0, 0, 1, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 1, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 1, 0, 0, 0, 1, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				1, 1, 1, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 1, 0,
			},
		},
		[]Node{
			{0, 0, "a", "", "nil", "myplaybook.yml", nil, nil, "", "", nil},
			{1, 0, "b", "", "shell", "myscript.sh", nil,
				map[string]string{
					"output1": "",
				},
				"", "", nil},
			{2, 0, "c", "", "shell", "myscript2.sh",
				[]string{
					"-e", "get_attribute 1:output1",
				}, nil, "", "", nil},
			{3, 0, "d", "", "nil", "myplaybook3.yml", nil, nil, "", "", nil},
			{4, 0, "e", "", "nil", "myplaybook4.yml", nil, nil, "", "", nil},
			{5, 0, "f", "", "nil", "myplaybook5.yml", nil, nil, "", "", nil},
			{6, 0, "g", "", "nil", "myplaybook6.yml", nil, nil, "", "", nil},
			{7, 0, "h", "", "nil", "myplaybook7.yml", nil, nil, "", "", nil},
		},
		nil,
		"",
	}
}
