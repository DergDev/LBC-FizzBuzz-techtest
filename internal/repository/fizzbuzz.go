package repository

import (
	"fmt"
	"techtest/internal/configs"
	"techtest/internal/models"
)

//Add or Update the current query to the DB
func AddQueryEntryToDB(url string) {
	db := configs.GetDBClient()

	var statistic models.Statistics
	db.Where("path = ?", url).Find(&statistic)
	if statistic.Path == "" {
		fmt.Println("Statistic created : " + url)
		db.Create(&models.Statistics{Path: url, Count: 1})
	} else {
		statistic.Count = statistic.Count + 1
		fmt.Println("Statistic updated : " + url + "count: " + fmt.Sprint(statistic.Count))
		db.Save(&statistic)
	}
}
