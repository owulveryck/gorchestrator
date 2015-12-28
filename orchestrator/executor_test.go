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
	"testing"
)

func TestInit(t *testing.T) {
	// OK test
	exe := ExecutorBackend{
		"https://localhost:8585/v1",
		"./test/cert.pem",
		"./test/key.pem",
		"./test/cert.pem",
		"/ping",
		nil,
	}
	err := exe.Init()
	if err != nil {
		t.Error(err)
	}
	exe = ExecutorBackend{
		"https://localhost:8585/v1",
		"./test/cert_nil.pem",
		"./test/key.pem",
		"./test/cert.pem",
		"/ping",
		nil,
	}
	err = exe.Init()
	if err == nil {
		t.Error(err)
	}
	exe = ExecutorBackend{
		"https://localhost:8585/v1",
		"./test/cert.pem",
		"./test/key.pem",
		"./test/cert_nil.pem",
		"/ping",
		nil,
	}
	err = exe.Init()
	if err == nil {
		t.Error(err)
	}
}
