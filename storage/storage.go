package storage

import (
	"fmt"
	"io"
	"log"
	"os"

	storage_go "github.com/supabase-community/storage-go"
)

var serviceKey string
var SupabaseURL string
var StorageClient *storage_go.Client

func InitStorage(){
	SupabaseURL = os.Getenv("supabaseURL")
	serviceKey = os.Getenv("supabaseAPIkey")
	endpoint := SupabaseURL+"/storage/v1"

	if SupabaseURL=="" || serviceKey ==""{
		log.Fatal("failed find key")
	}

	StorageClient= storage_go.NewClient(endpoint, serviceKey, nil)
	log.Println("bucket storage have been initialized")
}

func UploadFile(
	bucketName string,
	fileName string,
	file io.Reader,
	contentType string,
	)(string, error){
		if StorageClient == nil {
        log.Println("FATAL: StorageClient have not been initialized")
        return "", fmt.Errorf("storage client not initialized")
    }
	log.Printf("Storage: uploading file %s to bucket %s...", fileName, bucketName)


	upsertStatus := false
	_, err := StorageClient.UploadFile(
        bucketName,
        fileName,
        file,
        storage_go.FileOptions{
            ContentType: &contentType,
            Upsert:      &upsertStatus,
        },
	)
	if err != nil {
        log.Printf("Storage: failed to upload file %v", err)
        return "", err
    }

    log.Printf("Storage: file %s have been uploaded succesfuly.", fileName)
    publicURL := SupabaseURL + "/storage/v1/object/public/" + bucketName + "/" + fileName
    return publicURL, nil
}