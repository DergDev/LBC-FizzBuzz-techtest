package models

import "gorm.io/gorm"

type Statistics struct {
	gorm.Model
	Path  string
	Count int
}

type StatisticsResponse struct {
	Content []string `json:"content"`
	Hits    int      `json:"hits"`
}
