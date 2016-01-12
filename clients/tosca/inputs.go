package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Inputs map[string]Value

type Value struct {
	Value string `yaml:"value"`
}

// getInputs reade file f and returns its inputs
func getInputs(f string) (Inputs, error) {
	var i Inputs

	data, err := ioutil.ReadFile(f)
	if err != nil {
		return i, err
	}

	err = yaml.Unmarshal(data, &i)
	if err != nil {
		return i, err
	}
	return i, nil
}
