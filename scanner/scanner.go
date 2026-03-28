package scanner

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func HasIndex(path string, searchIndex string) (int, bool) {
	file, err := os.Open(path)
	if err != nil {
		return 0, false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lowerCaseSearchIndex := strings.ToLower(searchIndex)

	lineCount := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), lowerCaseSearchIndex) {
			return lineCount, true
		}
		lineCount++
	}

	return 0, false
}

func StreamFilesWithIndex(path string, searchIndex string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Erro ao ler diretório: ", err)
		return
	}

	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			wg.Add(1)
			go StreamFilesWithIndex(fullPath, searchIndex, results, wg)
		} else if entry.Type().IsRegular() {
			absolutePath, err := filepath.Abs(fullPath)
			if err != nil {
				fmt.Println("Erro ao ler caminho absoluto: ", err)
				return
			}

			line, hasIndex := HasIndex(absolutePath, searchIndex)

			if hasIndex {
				results <- fmt.Sprintf("%s:%d", absolutePath, line)
			}
		}
	}
}

func ScanFilesForIndex(path string, searchIndex string) []string {
	fmt.Println("Lendo diretório atual...")
	var wg sync.WaitGroup

	results := make(chan string)

	wg.Add(1)

	go StreamFilesWithIndex(path, searchIndex, results, &wg)

	go func() {
		wg.Wait()
		close(results)
	}()

	var foundPaths []string

	for path := range results {
		foundPaths = append(foundPaths, path)
	}

	return foundPaths
}
