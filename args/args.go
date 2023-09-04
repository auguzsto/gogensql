package args

import (
	"gogensql/handler"
	"os"
)

var Args []string

func GetColumnDateInArg() string {
	Args = os.Args[1:]
	data := Args[0]

	handler.NoArguments(Args)

	return data
}

func GetColumnOrgamentoInArg() string {
	Args = os.Args[1:]
	orcamento := Args[1]

	handler.NoArguments(Args)

	return orcamento
}

func GetColumnMargemInArg() string {
	Args = os.Args[1:]
	margem := Args[2]

	handler.NoArguments(Args)

	return margem
}

func GetFilePathInArg() string {
	Args = os.Args[1:]
	handler.NoArguments(Args)

	file := Args[3]

	return file
}
