package main

import (
	"encoding/json"
	"io/ioutil"
)

type Response struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	HasBeach    bool    `json:"hasBeach"`
	HasMountain string  `json:"hasMountain"`
	TempC       float64 `json:"tempC"`
}

type DataReader interface {
	ReadData() ([]Response, error)
}

type reader struct {
	path string
}

// NewReader initialises a new reader
func NewReader() DataReader {
	return &reader{
		path: "cities.json",
	}
}

func (r *reader) ReadData() ([]Response, error) {
	file, err := ioutil.ReadFile(r.path)
	if err != nil {
		return nil, err
	}
	var data []Response
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
