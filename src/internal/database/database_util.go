package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"onefor.fun/gosmarty/pkg/config"
)

func RecordUploadInDatabase(cfg *config.Config, fileName string, fileSize int64) error {
	connString := fmt.Sprintf("postgresql://%s:%s@%s/%s",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBName)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "INSERT INTO uploaded_files (file_name, file_size) VALUES ($1, $2)", fileName, fileSize)
	if err != nil {
		return err
	}

	return nil
}
