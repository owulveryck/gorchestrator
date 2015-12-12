/*
Olivier Wulveryck - author of Gexecutor
Copyright (C) 2015 Olivier Wulveryck

This file is part of the Gexecutor project and
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

package executor

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
)

type Task struct {
	ID string `json:"id"`
}

func uuid() Task {
	var t Task
	u := make([]byte, 16)
	_, err := rand.Read(u)
	if err != nil {
		return t
	}
	u[8] = (u[8] | 0x80) & 0xBF // what does this do?
	u[6] = (u[6] | 0x40) & 0x4F // what does this do?
	t.ID = hex.EncodeToString(u)
	return t
}

// This will hold all the requested tasks
var tasks map[string]node

func Run() {

	tasks = make(map[string]node, 0)
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
