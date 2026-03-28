package main

import (
	"fmt"
	"os"
	"super-grep/db"
	"super-grep/scanner"
)

func main() {
	args := os.Args
	if len(args) == 0 {
		panic("Erro: termo de pesquisa inválido!")
	}

	searchIndex := os.Args[1]

	db.InitDB()

	history, err := db.FindTerm(searchIndex)
	if err == nil && len(history) > 0 {
		fmt.Printf("O termo já existe no banco com %d ocorrencias passsadas.", len(history))
	}

	fmt.Printf("Iniciando busca paralela por %s\n", searchIndex)
	foundPaths := scanner.ScanFilesForIndex(".", searchIndex)

	fmt.Printf("Busca conluida, total de ocorrências encontradas agora: %d\n", len(foundPaths))

	if len(foundPaths) > 0 {
		err := db.UpsertPath(searchIndex, foundPaths)
		if err != nil {
			fmt.Println("Erro ao escrever no banco de dados: ", err)
			return
		}
		fmt.Println("Índice do banco de dados atualizado com sucesso.")
	} else {
		fmt.Println("Nenhuma ocorrência nova para salvar no índice.")
	}

	fmt.Println("Resultados:")
	for _, path := range foundPaths {
		fmt.Println(path)
	}
}
