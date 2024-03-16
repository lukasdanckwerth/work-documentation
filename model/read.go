package model

import (
	"encoding/json"
	"os"
)

func ReadDay(thePath string) *WorkDay {

	b, err := os.ReadFile(thePath)

	if err != nil {
		return nil
	}

	var day *WorkDay
	err = json.Unmarshal([]byte(b), &day)

	if err != nil {
		panic(err)
	}

	return day
}

func WriteDay(day *WorkDay, thePath string) error {
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
