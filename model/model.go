package model

import (
	"fmt"
	"slices"
)

type Day struct {
	Date    int64   `json:"date"`
	Entries []Entry `json:"entries"`
}

func (d *Day) AddEntry(entry Entry) {
	d.Entries = slices.Insert(d.Entries, len(d.Entries), entry)
}

type Entry struct {
	Title   string `json:"title"`
	Start   int64  `json:"start"`
	Created int64  `json:"created"`
	Length  int64  `json:"length"`
}

func (e *Entry) LenghtFormatted() string {
	return fmt.Sprintf("%02d:%02d", e.Length/60, e.Length%60)
}
