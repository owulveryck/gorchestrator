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

package main

import (
	"flag"
)

var toscaFilename string
var inputFilename string

func init() {
	const (
		defaultTosca = ""
		usageTosca   = "A TOSCA compliant Yaml file containg a complete Service Template Definition"
		defaultInput = ""
		usageInput   = "An Input file in yaml that handles the inputs values corresponding to the TOSCA file"
	)
	flag.StringVar(&toscaFilename, "tosca", defaultTosca, usageTosca)
	flag.StringVar(&toscaFilename, "t", defaultTosca, usageTosca+" (shorthand)")
	flag.StringVar(&inputFilename, "input", defaultInput, usageInput)
	flag.StringVar(&inputFilename, "i", defaultInput, usageInput+" (shorthand)")

}
