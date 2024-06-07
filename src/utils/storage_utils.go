package utils

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"github.com/jackc/pgx/v4"
)

func UploadToCloudStorage(objectName string, file io.Reader) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	bucketName := os.Getenv("CLOUD_STORAGE_BUCKET")
	bucket := client.Bucket(bucketName)
	obj := bucket.Object(objectName)

	// Upload the file to Cloud Storage
	wc := obj.NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}

func RecordUploadInDatabase(fileName string, fileSize int64) error {
	connString := fmt.Sprintf("postgresql://%s:%s@%s/%s",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

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
