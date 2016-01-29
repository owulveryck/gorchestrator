package structure

import (
	"fmt"
	"math"
	"sync"
)

// Matrix is a list representation of a squared matrix
// it concurrent safe
type Matrix struct {
	Matrix []int64
	sync.RWMutex
}

// isValid check if the matrix is squared
func (m *Matrix) isValid() error {
	l := math.Sqrt(float64(len((*m).Matrix)))
	if float64(int64(l)) != l {
		return fmt.Errorf("Matrix is not a squared one")
	}
	return nil
}

// Dim returns the dimension of the matrix
func (m *Matrix) Dim() int {
	m.RLock()
	defer m.RUnlock()
	err := m.isValid()
	if err != nil {
		return 0
	}
	return int(math.Sqrt(float64(len((*m).Matrix))))
}

// Get sets the value v in row r and column c
func (m *Matrix) Set(r, c int, v int64) {
	m.Lock()
	defer m.Unlock()
	i := m.Dim()
	(*m).Matrix[r*i+c] = v
}

// Get returns the value in row r and column c
func (m *Matrix) At(r, c int) int64 {
	m.RLock()
	defer m.RUnlock()
	i := m.Dim()
	return (*m).Matrix[r*i+c]
}

func (m *Matrix) Sum() int64 {
	m.RLock()
	defer m.RUnlock()
	var v int64
	for r := 0; r < m.Dim(); r++ {
		for c := 0; c < m.Dim(); c++ {
			v = v + m.At(r, c)
		}
	}
	return v
}
