package controller

import (
	"fmt"
	"gogensql/model"
)

func GenSQLReturn(meta model.Excel) string {
	sql := fmt.Sprintf(`update pcmeta set vlvendaprev = %v,mixprev = %v where data BETWEEN '%v' and '%v' and codfilial = 6 and codusur = 100 and codigo = %v and tipometa = 'S';`, meta.Orcamento, meta.Margem, meta.DataInicial, meta.DataFinal, meta.Secao)

	return sql
}
