package main

import (
	"encoding/json"
	"os"
)

func LoadState(filename string) (State, error) {
	var s State
	data, err := os.ReadFile(filename)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(data, &s)
	return s, err
}

func SaveState(filename string, s State) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0o644)
}
