package main

import (
	"fmt"
	"gogensql/args"
	"gogensql/controller"
	"gogensql/model"

	"github.com/xuri/excelize/v2"
)

// Order args: COLUMN-DATA - COLUMN-ORCAMENTO - COLUMN-MARGEM - PATH_FILE
// Exemple: ./gogesql M4 M O ./file.xlsx

func main() {
	excel, err := excelize.OpenFile(args.GetFilePathInArg())
	if err != nil {
		panic(err)
	}
	defer excel.Close()

	sheet := "Plan1"
	rows, err := excel.GetRows(sheet)
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		i = i + 1

		if i >= 7 {

			meta := model.Excel{
				Secao:       row[0],
				Orcamento:   controller.GetOrcamento(i, excel, sheet),
				Margem:      controller.GetMargem(i, excel, sheet),
				DataInicial: controller.GetPeriodicData(excel, "init", sheet),
				DataFinal:   controller.GetPeriodicData(excel, "final", sheet),
			}

			fmt.Println(controller.GenSQLReturn(meta))
		}
	}
}
