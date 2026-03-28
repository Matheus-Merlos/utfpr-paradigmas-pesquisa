package main

import (
	"fmt"
	"os"
)

func ListFolderContent(path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Erro ao ler diretório: ", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			ListFolderContent(entry.Name())
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Println("Ocorreu um erro ao ler as informações do arquivo: ", err)
				return
			}
			fmt.Println("Arquivo: ", info.Name())
		}
	}
}

func main() {
	ListFolderContent(".")
}
