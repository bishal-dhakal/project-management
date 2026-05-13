package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/jmoiron/sqlx"
)

func RunMigrations(db *sqlx.DB, migrationsDir string) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			filename VARCHAR(255) UNIQUE NOT NULL,
			applied_at TIMESTAMP DEFAULT NOW()
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))
	if err != nil {
		return err
	}
	sort.Strings(files)

	for _, file := range files {
		filename := filepath.Base(file)

		var count int
		db.QueryRow(`SELECT COUNT(*) FROM migrations WHERE filename = $1`, filename).Scan(&count)

		if count > 0 {
			fmt.Printf("skipping migration: %s\n", filename)
			continue
		}

		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", filename, err)
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("failed to apply %s: %w", filename, err)
		}
		db.Exec(`INSERT INTO migrations (filename) VALUES ($1)`, filename)
		fmt.Printf("applied migration: %s\n", filename)
	}
	return nil
}
