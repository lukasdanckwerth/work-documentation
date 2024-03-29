package model

import (
	"fmt"
	"slices"
	"time"
)

// A work day entity
type WorkDay struct {
	Date    int64   `json:"date"`
	Entries []Entry `json:"entries"`
}

// Add a new entry to the work day
func (d *WorkDay) AddEntry(entry Entry) {
	d.Entries = slices.Insert(d.Entries, len(d.Entries), entry)
}

// An entry in a work day
type Entry struct {
	Title   string `json:"title"`
	Created int64  `json:"created"`
	Length  int64  `json:"length"`
	UUID    string `json:"uuid"`
}

// Returns a formatted version of the lenght of the entry
func (e *Entry) LenghtFormatted() string {
	return fmt.Sprintf("%02d:%02d", e.Length/60, e.Length%60)
}

func (e *Entry) CreatedFormatted() string {
	return time.UnixMilli(e.Created).Format("2006-01-02 15:04")
}

func (e *Entry) CreatedTimeFormatted() string {
	return time.UnixMilli(e.Created).Format("15:04")
}
