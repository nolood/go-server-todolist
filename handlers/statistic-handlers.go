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

func getYears() []int {
	currentYear := time.Now().Year()
	years := make([]int, 6)
	for i := 0; i < 6; i++ {
		years[i] = currentYear - 5 + i
	}
	return years
}

func getMonths() []string {
	return []string{
		"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь",
		"Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь",
	}
}

func getLabels(period string, years []int, months []string) []string {
	switch period {
	case "year":
		stringsYears := make([]string, 0, 6)
		for _, num := range years {
			stringsYears = append(stringsYears, strconv.Itoa(num))
		}
		return stringsYears
	case "month":
		return months
	default:
		return make([]string, 0)
	}
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

	months := getMonths()

	years := getYears()

	type Dataset struct {
		Label           string `json:"label"`
		Data            []int  `json:"data"`
		BackgroundColor string `json:"backgroundColor"`
	}

	groupedRecords := make(map[string]map[int]int)
	colors := make(map[string]string)
	var income, expense, profit, loss map[int]int

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

		if params.Type == 0 {
			if record.RecordTypeID == 1 {
				if income == nil {
					income = make(map[int]int)
				}
				income[unit] += record.Amount
			} else if record.RecordTypeID == 2 {
				if expense == nil {
					expense = make(map[int]int)
				}
				expense[unit] += record.Amount
			}
		}
	}

	if params.Type == 0 {
		profit = make(map[int]int)
		loss = make(map[int]int)
		for unit, incomeAmount := range income {
			expenseAmount := expense[unit]

			profitAmount := incomeAmount - expenseAmount
			lossAmount := expenseAmount - incomeAmount

			if lossAmount > 0 {
				profitAmount = 0
			} else {
				lossAmount = 0
			}

			profit[unit] = profitAmount
			loss[unit] = lossAmount
		}

		groupedRecords = make(map[string]map[int]int)
		groupedRecords["Доход"] = income
		groupedRecords["Расход"] = expense
		groupedRecords["Прибыль"] = profit
		groupedRecords["Убыток"] = loss
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

		var backgroundColor string
		switch title {
		case "Доход":
			backgroundColor = "green.500"
		case "Расход":
			backgroundColor = "red.500"
		case "Прибыль":
			backgroundColor = "blue.500"
		case "Убыток":
			backgroundColor = "orange.500"
		default:
			backgroundColor = colors[title]
		}

		dataset := Dataset{
			Label:           title,
			Data:            amounts,
			BackgroundColor: backgroundColor,
		}
		datasets = append(datasets, dataset)
	}

	type Data struct {
		Labels   interface{} `json:"labels"`
		Datasets []Dataset   `json:"datasets"`
	}

	labels := getLabels(params.Period, years, months)

	data := Data{
		Labels:   labels,
		Datasets: datasets,
	}

	w.Write(toJson(data))
}
