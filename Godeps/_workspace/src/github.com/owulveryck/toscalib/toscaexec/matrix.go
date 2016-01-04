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

package toscaexec

import (
	"fmt"
	"math"
)

// Matrix is a list representation of a squared matrix
type Matrix []int64

// isValid check if the matrix is squared
func (m *Matrix) isValid() error {
	l := math.Sqrt(float64(len(*m)))
	if float64(int64(l)) != l {
		return fmt.Errorf("Matrix is not a squared one")
	}
	return nil
}

// Dim returns the dimension of the matrix
func (m *Matrix) Dim() int {
	err := m.isValid()
	if err != nil {
		return 0
	}
	return int(math.Sqrt(float64(len(*m))))
}

// Get sets the value v in row r and column c
func (m *Matrix) Set(r, c int, v int64) {
	i := m.Dim()
	(*m)[r*i+c] = v
}

// Get returns the value in row r and column c
func (m *Matrix) At(r, c int) int64 {
	i := m.Dim()
	return (*m)[r*i+c]
}

func (m *Matrix) Sum() int64 {
	var v int64
	for r := 0; r < m.Dim(); r++ {
		for c := 0; c < m.Dim(); c++ {
			v = v + m.At(r, c)
		}
	}
	return v
}
