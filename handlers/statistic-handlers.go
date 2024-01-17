package handlers

import (
	"go-server/internal/config"
	"go-server/internal/storage/postgres"
	"net/http"
	"strconv"
	"time"
)

type GetStatisticParams struct {
	BillID int    `json:"bill_id"`
	Type   int    `json:"type"`
	Period string `json:"period"`
}

func parseYear(date string) int {
	year, err := strconv.Atoi(date[:4])
	if err != nil {
		config.Logger.Error(err.Error())
		return 0
	}
	return year
}

func parseMonth(date string) int {
	month, err := strconv.Atoi(date[5:7])
	if err != nil {
		config.Logger.Error(err.Error())
		return 0
	}
	return month
}

func GetStatistic(w http.ResponseWriter, r *http.Request) {
	var params GetStatisticParams

	err := validateQueryParams(r, &params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := getUserId(r)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var bill postgres.Bill

	query := postgres.Db.Table("bills")

	query = query.Where("user_id = ?", userId).Where("id = ?", params.BillID)

	query.Find(&bill)

	if query.Error != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var records []postgres.Record

	query = postgres.Db.Table("records")

	query = query.Where("bill_id = ?", params.BillID)

	if params.Type != 0 {
		query = query.Where("record_type_id = ?", params.Type)
	}

	query = query.Preload("Article")

	query.Find(&records)

	years := make([]int, 6)
	months := make([]string, 12)

	if params.Period == "year" {
		currentYear := time.Now().Year()
		for i := 0; i < 6; i++ {
			years[i] = currentYear - 5 + i
		}
	}

	if params.Period == "month" {
		months = []string{
			"Январь",
			"Февраль",
			"Март",
			"Арпель",
			"Май",
			"Июнь",
			"Июль",
			"Август",
			"Сентябрь",
			"Октбярь",
			"Ноябрь",
			"Декабрь",
		}
	}

	type Dataset struct {
		Label           string `json:"label"`
		Data            []int  `json:"data"`
		BackgroundColor string `json:"backgroundColor"`
	}

	groupedRecords := make(map[string]map[int]int)
	colors := make(map[string]string)

	for _, record := range records {
		title := record.Article.Title
		var unit int
		switch params.Period {
		case "year":
			unit = parseYear(record.Date)
		case "month":
			unit = parseMonth(record.Date)
		default:
			unit = 0
		}

		if _, ok := groupedRecords[title]; !ok {
			groupedRecords[title] = make(map[int]int)
			colors[title] = record.Article.Color
		}

		groupedRecords[title][unit] += record.Amount
	}

	var datasets []Dataset
	for title, amountsByUnit := range groupedRecords {
		var units []int
		switch params.Period {
		case "year":
			units = years
		case "month":
			units = make([]int, len(months))
			for i := range months {
				units[i] = i + 1
			}
		}

		amounts := make([]int, len(units))
		for i, unit := range units {
			amounts[i] = amountsByUnit[unit]
		}

		dataset := Dataset{
			Label:           title,
			Data:            amounts,
			BackgroundColor: colors[title],
		}
		datasets = append(datasets, dataset)
	}

	type Data struct {
		Labels   interface{} `json:"labels"`
		Datasets []Dataset   `json:"datasets"`
	}

	data := Data{
		Datasets: datasets,
	}

	if params.Period == "year" {
		data.Labels = years
	}

	if params.Period == "month" {
		data.Labels = months
	}

	w.Write(toJson(data))
}
