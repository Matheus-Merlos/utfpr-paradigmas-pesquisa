package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListFolderContent(path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Erro ao ler diretório: ", err)
		return
	}

	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			ListFolderContent(fullPath)
		} else {
			absolutePath, err := filepath.Abs(fullPath)
			if err != nil {
				fmt.Println("Erro ao ler caminho absoluto: ", err)
				return
			}
			fmt.Println(absolutePath)
		}
	}
}

func main() {
	ListFolderContent(".")
}
