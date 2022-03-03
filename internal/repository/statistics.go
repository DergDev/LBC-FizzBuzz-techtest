package repository

import (
	"errors"
	"strings"
	"techtest/internal/configs"
	"techtest/internal/models"

	"gorm.io/gorm"
)

//Get the most requested query from the DB
func GetQueryEntry() (models.StatisticsResponse, error) {
	db := configs.GetDBClient()
	var statistic models.Statistics
	err := db.Order("count desc").First(&statistic).Error
	statisticResponse := models.StatisticsResponse{}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return statisticResponse, err
	}
	statisticResponse.Content = strings.Split(statistic.Path, "&")
	statisticResponse.Hits = statistic.Count
	return statisticResponse, nil
}
