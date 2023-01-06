package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	return err
}

func main() {
	filename := "demo.txt"

	file, err := os.Create(filename)

	file.WriteString("Hello,welcome to AWS tutorial")

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" %s  \n", data)
	defer file.Close()

	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "AlekhyaGandra",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})
	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}
	//bucketName := "learnaws-go-sdk-tutorial-cloudquicklabs"
	uploader := s3manager.NewUploader(sess)
	//filename := "demo.txt"
	err = UploadFile(uploader, "demo.txt", "demobucket1s3", "demo.txt")
	if err != nil {
		fmt.Printf("Failed to upload file: %v", err)
	}
	fmt.Println("Successfully uploaded file!")

}
