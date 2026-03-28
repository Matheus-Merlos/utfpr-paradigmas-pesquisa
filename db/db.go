package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

type FoundResult struct {
	TermID           int
	Term             string
	AbsolutePathLine string
}

func InitDB() {

	dataDir := filepath.Join(os.Getenv("HOME"), ".local", "share", "termdb")

	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Fatal("Erro ao criar pasta do banco:", err)
	}

	dbPath := filepath.Join(dataDir, "termdb.db")

	var err error

	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal("Erro ao abrir o banco de dados:", err)
	}

	_, err = DB.Exec(`
        PRAGMA journal_mode = WAL;
        PRAGMA foreign_keys = ON;
    `)
	if err != nil {
		log.Println("Aviso: não foi possível aplicar pragmas:", err)
	}

	fmt.Println("Successfully connected to SQLite!")

	createTables()
}

func createTables() {
	query := `
	CREATE TABLE IF NOT EXISTS termList (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		term VARCHAR(255) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS foundDirs (
	    idTerm INTEGER NOT NULL,
		absolutePathLine TEXT NOT NULL,
		FOREIGN KEY (idTerm) REFERENCES termList(id) ON DELETE CASCADE
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Erro ao criar tabelas: %v", err)
	}

	fmt.Println("Table created (or already exists)")
}

func InsertPath(term string, absolutePaths []string) error {

	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var termID int64
	err = tx.QueryRow(`
        INSERT INTO termList (term) 
        VALUES (?) 
        ON CONFLICT(term) DO UPDATE SET term = excluded.term
        RETURNING id
    `, term).Scan(&termID)
	if err != nil {
		return fmt.Errorf("erro ao inserir termo: %w", err)
	}

	stmt, err := tx.Prepare(`
        INSERT INTO foundDirs (idTerm, absolutePathLine)
        VALUES (?, ?)
    `)

	if err != nil {
		return fmt.Errorf("erro ao preparar statement: %w", err)
	}

	defer stmt.Close()

	for _, path := range absolutePaths {
		if path == "" {
			continue
		}
		_, err = stmt.Exec(termID, path)
		if err != nil {
			return fmt.Errorf("erro ao inserir caminho '%s': %w", path, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("erro ao fazer commit: %w", err)
	}

	return nil
}

func findTerm(term string) ([]FoundResult, error) {
	query := `
        SELECT t.id, t.term, fd.absolutePathLine
        FROM termList t
        JOIN foundDirs fd ON t.id = fd.idTerm
        WHERE t.term LIKE ?
	`

	rows, err := DB.Query(query, "%"+term+"%")
	if err != nil {
		return nil, fmt.Errorf("erro na consulta: %w", err)
	}
	defer rows.Close()

	var results []FoundResult

	for rows.Next() {
		var r FoundResult
		if err := rows.Scan(&r.TermID, &r.Term, &r.AbsolutePathLine); err != nil {
			return nil, fmt.Errorf("erro ao ler resultado: %w", err)
		}
		results = append(results, r)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante iteração: %w", err)
	}

	return results, nil
}
