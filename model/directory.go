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
		Path: dirname + "/.wodo/data",
	}
}

func (d *Directory) GetDirAndPath(year string, month string, day string) (string, string) {
	theDir := strings.Join([]string{d.Path, year, month}, "/")
	thePath := strings.Join([]string{theDir, day}, "/") + ".json"
	return theDir, thePath
}

// Receives the WorkDay
func (d *Directory) ReceiveWorkday(year string, month string, day string) *WorkDay {

	theDir, thePath := d.GetDirAndPath(year, month, day)

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

// Write the given WorkDay
func (d *Directory) WriteWorkday(dayObj *WorkDay, year string, month string, day string) {
	theDir, thePath := d.GetDirAndPath(year, month, day)

	os.MkdirAll(theDir, os.ModePerm)

	err := WriteDay(dayObj, thePath)

	if err != nil {
		panic(err)
	}
}

// Reveive a ann array of WorkDays for a given month and year
func (d *Directory) ReceiveMonth(year string, month string) []*WorkDay {
	theDir, _ := d.GetDirAndPath(year, month, "")
	entries, err := os.ReadDir(theDir)

	// simply return an empty array if the directory doesn't exist
	if err != nil {
		return []*WorkDay{}
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
