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

type lock int

func (l *lock) Lock() {
	*l = 1
}
func (l *lock) Unlock() {
	*l = 0
}

func fanOut(outputs ...chan<- Graph) chan<- Graph {
	c := make(chan Graph)
	for i := range outputs {
		output := outputs[i] // New instance of 'input' for each loop.
		go func(i int) {
			for {
				output <- <-c
			}
		}(i)
	}
	return c
}
func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i] // New instance of 'input' for each loop.
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}
