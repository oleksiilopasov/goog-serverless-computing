package storage

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
)

func UploadToCloudStorage(bucketName, objectName string, file io.Reader) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

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
