package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"onefor.fun/gosmarty/pkg/config"
)

// EnsureTableExists ensures that the uploaded_files table exists.
func EnsureTableExists(conn *pgx.Conn) error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS uploaded_files (
		id SERIAL PRIMARY KEY,
		file_name TEXT NOT NULL,
		file_size BIGINT NOT NULL,
		uploaded_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := conn.Exec(context.Background(), createTableQuery)
	return err
}

func RecordUploadInDatabase(cfg *config.Config, fileName string, fileSize int64) error {
	connString := fmt.Sprintf("postgresql://%s:%s@%s/%s",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBName)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	// Ensure the uploaded_files table exists
	if err := EnsureTableExists(conn); err != nil {
		return fmt.Errorf("error ensuring table exists: %w", err)
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO uploaded_files (file_name, file_size) VALUES ($1, $2)", fileName, fileSize)
	if err != nil {
		return err
	}

	return nil
}
