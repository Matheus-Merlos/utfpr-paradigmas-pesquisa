package main

import (
	"fmt"
	"os"
	"super-grep/scanner"
)

func main() {
	args := os.Args
	if len(args) == 0 {
		panic("Erro: termo de pesquisa inválido!")
	}

	foundPaths := scanner.ScanFilesForIndex(".", os.Args[1])

	fmt.Printf("Busca conluida, total de pastas encontradas: %d\n", len(foundPaths))

	for _, path := range foundPaths {
		fmt.Println(path)
	}
}
