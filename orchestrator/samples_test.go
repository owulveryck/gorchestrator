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
	"time"
)

var valid Graph
var validAndNoArtifact Graph
var validAndFailed Graph
var validAndSleep Graph
var validAndSleepAndFailed Graph
var notValid Graph

func init() {
	validAndNoArtifact = Graph{"ValidAndNoArtifact",
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
			{0, 0, "a", "nil", "", nil, nil},
			{1, 0, "b", "", "nil", nil, nil},
			{2, 0, "c", "", "", nil, nil},
			{3, 0, "d", "", "", nil, nil},
			{4, 0, "e", "", "", nil, nil},
			{5, 0, "f", "", "", nil, nil},
			{6, 0, "g", "", "", nil, nil},
			{7, 0, "h", "", "", nil, nil},
		},
		time.After(30 * time.Second),
	}

	validAndSleep = Graph{"ValidAndSleep",
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
			{0, 0, "a", "sleep", "", nil, nil},
			{1, 0, "b", "sleep", "", nil, nil},
			{2, 0, "c", "sleep", "", nil, nil},
			{3, 0, "d", "sleep", "", nil, nil},
			{4, 0, "e", "sleep", "", nil, nil},
			{5, 0, "f", "sleep", "", nil, nil},
			{6, 0, "g", "sleep", "", nil, nil},
			{7, 0, "h", "sleep", "", nil, nil},
		},
		time.After(30 * time.Second),
	}

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
			{0, 0, "a", "nil", "myplaybook.yml", nil, nil},
			{1, 0, "b", "nil", "myscript.sh", nil,
				map[string]string{
					"result": "",
				},
			},
			{2, 0, "c", "nil", "myscript2.sh",
				[]string{
					"-e", "get_attribute b:result",
				}, nil},
			{3, 0, "d", "nil", "myplaybook3.yml", nil, nil},
			{4, 0, "e", "nil", "myplaybook4.yml", nil, nil},
			{5, 0, "f", "nil", "myplaybook5.yml", nil, nil},
			{6, 0, "g", "nil", "myplaybook6.yml", nil, nil},
			{7, 0, "h", "nil", "myplaybook7.yml", nil, nil},
		},
		time.After(30 * time.Second),
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
			{0, 0, "a", "nil", "myplaybook.yml", nil, nil},
			{1, 0, "b", "shell", "myscript.sh", nil,
				map[string]string{
					"output1": "",
				},
			},
			{2, 0, "c", "shell", "myscript2.sh",
				[]string{
					"-e", "get_attribute 1:output1",
				}, nil},
			{3, 0, "d", "nil", "myplaybook3.yml", nil, nil},
			{4, 0, "e", "nil", "myplaybook4.yml", nil, nil},
			{5, 0, "f", "nil", "myplaybook5.yml", nil, nil},
			{6, 0, "g", "nil", "myplaybook6.yml", nil, nil},
			{7, 0, "h", "nil", "myplaybook7.yml", nil, nil},
		},
		nil,
	}
}
