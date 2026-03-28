package main

import (
	"fmt"
	"os"
	"super-grep/scanner"
	db "super-grep/storage"
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
		fmt.Printf("O termo já existe no banco com %d ocorrencias passsadas nos seguintes diretórios: \n", len(history))
		for i, res := range history {
			fmt.Printf(" [%02d] %s\n", i+1, res.AbsolutePathLine)
		}
	}

	var choice string
	var foundPaths []string
	if len(history) > 0 {
		fmt.Println("Você gostaria de buscar novamente?")
		fmt.Print("Sua escolha [y/n(default)]: ")
		fmt.Scanln(&choice)
	}

	if choice == "y" || len(history) == 0 {
		fmt.Printf("Iniciando busca paralela por %s\n", searchIndex)
		foundPaths = scanner.ScanFilesForIndex(".", searchIndex)

		fmt.Printf("Busca concluida, total de ocorrências encontradas agora: %d\n", len(foundPaths))

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
	} else {
		fmt.Println("Carregando resultados do banco de dados...")
		for _, item := range history {
			foundPaths = append(foundPaths, item.AbsolutePathLine)
		}
	}

	fmt.Println("Resultados:")
	for _, path := range foundPaths {
		fmt.Println(path)
	}
}
