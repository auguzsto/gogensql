package controller

import (
	"gogensql/args"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

func GetPeriodicData(excel *excelize.File, periodic string, sheet string) string {
	dataTime := getDataAndStyleDataExcel(excel, sheet)
	var data string
	switch periodic {
	case "init":
		data = formaterData(dataTime)
	case "final":
		data = formaterData(dataTime.AddDate(0, 1, -1))
	}

	return data
}

func GetOrcamento(iterable int, excel *excelize.File, sheet string) string {
	orcamento, err := excel.GetCellValue(sheet, args.GetColumnOrgamentoInArg()+strconv.Itoa(iterable))
	if err != nil {
		panic(err)
	}

	return orcamento
}

func GetMargem(iterable int, excel *excelize.File, sheet string) string {
	margem, err := excel.GetCellValue(sheet, args.GetColumnMargemInArg()+strconv.Itoa(iterable))
	if err != nil {
		panic(err)
	}

	return margem
}

func formaterData(data time.Time) string {
	day := strings.Split(data.Format(time.RFC850), " ")[1][:3]
	mouth := strings.Split(data.Format(time.RFC850), " ")[1][3:6]
	year := "-20" + strings.Split(data.Format(time.RFC850), " ")[1][7:]

	return day + mouth + year

}

func getDataAndStyleDataExcel(excel *excelize.File, sheet string) time.Time {
	styleData, _ := excel.NewStyle(&excelize.Style{
		NumFmt: 15,
	})

	excel.SetCellStyle(sheet, args.GetColumnDateInArg(), args.GetColumnDateInArg(), styleData)
	data, _ := excel.GetCellValue(sheet, args.GetColumnDateInArg())

	dataFormatedTime, err := time.Parse("1-Jan-06", data)
	if err != nil {
		panic(err)
	}

	return dataFormatedTime
}
