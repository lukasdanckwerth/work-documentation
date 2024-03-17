package model

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"
)

type Directory struct {
	Path string
}

// Creates and receive the working directory.
func WorkDirectory() *Directory {
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return &Directory{
		Path: dirname + "/.wodo",
	}
}

// Receives the WorkDay
func (d *Directory) ReceiveWorkday(year string, month string, day string) *WorkDay {
	theDir := strings.Join([]string{d.Path, year, month}, "/")
	thePath := strings.Join([]string{theDir, day}, "/") + ".json"

	os.MkdirAll(theDir, os.ModePerm)

	dayObj := ReadDay(thePath)

	if dayObj == nil {
		dayObj = &WorkDay{
			Date:    time.Now().UnixMilli(),
			Entries: []Entry{},
		}
	}

	return dayObj
}

func (d *Directory) WriteWorkday(dayObj *WorkDay, year string, month string, day string) {
	theDir := strings.Join([]string{d.Path, year, month}, "/")
	thePath := strings.Join([]string{theDir, day}, "/") + ".json"

	os.MkdirAll(theDir, os.ModePerm)

	err := WriteDay(dayObj, thePath)

	if err != nil {
		panic(err)
	}
}

func (d *Directory) ReceiveMonth(year string, month string) []*WorkDay {
	theDir := strings.Join([]string{d.Path, year, month}, "/")
	entries, err := os.ReadDir(theDir)

	if err != nil {
		panic(err)
	}

	days := []*WorkDay{}

	for _, entry := range entries {
		dayObj := ReadDay(strings.Join([]string{theDir, entry.Name()}, "/"))
		days = append(days, dayObj)
	}

	return days
}

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
