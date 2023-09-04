package handler

import "log"

func NoArguments(args []string) {
	//Data
	if len(args) <= 0 {
		log.Fatalln("É necessário passar, exatamente, a coluna que contêm a data. \n Exemplo: gogensql D4 ORCA MARGEM ./arquivo.xlsx")
	}

	//Orçamento
	if len(args) == 1 {
		log.Fatalln("É necessário passar a coluna que contêm os dados do orçamento. \n Exemplo: gogensql DATA O MARGEM ./arquivo.xlsx.")
	}

	//Margem
	if len(args) == 2 {
		log.Fatalln("É necessário passar a coluna que contêm os dados da margem. \n Exemplo: gogensql DATA ORCA M ./arquivo.xlsx")
	}

	//Arquivo
	if len(args) == 3 {
		log.Fatalln("É necessário passar o caminho do arquivo. \n Exemplo: gogensql data orca margem ./arquivo.xlsx")
	}
}
