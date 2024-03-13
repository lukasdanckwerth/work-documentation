package model

import "fmt"

func Format(length int64) string {
	return fmt.Sprintf("%02d:%02d", length/60, length%60)
}

func FormatCats(length int64) string {
	return fmt.Sprintf("%.2f", float64(length/60)+float64(length%60)/60)
}
