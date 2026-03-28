package main

import (
	"fmt"
	"os"
	"utfpr-paradigmas-pesquisa/scanner"
)

func main() {
	foundPaths := scanner.ScanFilesForIndex(".", os.Args[1])

	fmt.Printf("Busca conluida, total de pastas encontradas: %d\n", len(foundPaths))

	for _, path := range foundPaths {
		fmt.Println(path)
	}
}
