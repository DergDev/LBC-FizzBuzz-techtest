package repository

import (
	"techtest/internal/configs"
	"techtest/internal/models"
)

//Add or Update the current query to the DB
func AddQueryEntryToDB(url string) {
	db := configs.GetDBClient()

	var statistic models.Statistics
	db.Where("path = ?", url).Find(&statistic)
	if statistic.Path == "" {
		db.Create(&models.Statistics{Path: url, Count: 1})
	} else {
		statistic.Count = statistic.Count + 1
		db.Save(&statistic)
	}
}
