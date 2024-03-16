package model

import (
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
func (d *Directory) ReceiveDay(year string, month string, day string) *WorkDay {
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
