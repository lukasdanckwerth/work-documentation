package model

import (
	"encoding/json"
	"os"
)

func ReadDay(thePath string) *Day {

	b, err := os.ReadFile(thePath)

	if err != nil {
		return nil
	}

	var day *Day
	err = json.Unmarshal([]byte(b), &day)

	if err != nil {
		panic(err)
	}

	return day
}

func WriteDay(day *Day, thePath string) error {
	jsonBytes, err := json.MarshalIndent(day, "", " ")

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(thePath, jsonBytes, 0666)

	if err != nil {
		panic(err)
	}

	return nil
}
