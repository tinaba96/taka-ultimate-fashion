package s3Upload

import (
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadToS3(productId int) {
	log.Println("Start Uploading to S3")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	
	targetFilePath := "scraper/downloads/product_image_" + strconv.Itoa(productId) +".jpg"
	file, err := os.Open(targetFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bucketName := "tuf-products-data"
	objectKey := "products-images/product_image_" + strconv.Itoa(productId) +".jpg"
	awsRegion := "us-east-1"

    // create S3 client
	svc := s3.New(sess, &aws.Config{
		Region: aws.String(awsRegion),
	})

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key: aws.String(objectKey),
		Body: file,
	}

    // upload to S3
	_, err = svc.PutObject(params)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("S3 done! You are AMAZING!!")
    errDelete := os.Remove(targetFilePath)
	if errDelete != nil {
		log.Fatal(err)
	}
}
