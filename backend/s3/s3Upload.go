package s3Upload

import (
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// var (
//     client *s3.Client
// )

// type S3ObjectAPI interface {
//     PutObject(
//         ctx context.Context,
//         params *s3.PutObjectInput,
//         optFns ...func(*s3.Options),
//     ) (*s3.PutObjectOutput, error)
// }


func UploadToS3(productId int) {
	log.Println("Start Uploading to S3")

	// sess, _ := session.NewSession(&aws.Config{
	// 	Region: aws.String("us-west-1")},
	// )
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		// Profile:           "di",
		SharedConfigState: session.SharedConfigEnable,
		// Config: aws.Config{
		// 	Region: aws.String("us-west-1"),
		// },
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

	// Uploaderを作成し、ローカルファイルをアップロード
	// uploader := s3manager.NewUploader(sess)
	// _, err = uploader.Upload(&s3manager.UploadInput{
	// 	Bucket: aws.String(bucketName),
	// 	Key:    aws.String(objectKey),
	// 	Body:   file,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	// if err := s3.Put(bucketName, objectKey, file); err != nil {
	// 	panic(err)
	// }

	// S3クライアントを作成します
	svc := s3.New(sess, &aws.Config{
		Region: aws.String(awsRegion),
	})

	// S3にアップロードする内容をparamsに入れます
	params := &s3.PutObjectInput{
		// Bucket アップロード先のS3のバケット
		Bucket: aws.String(bucketName),
		// Key アップロードする際のオブジェクト名
		Key: aws.String(objectKey),
		// Body アップロードする画像ファイル
		Body: file,
	}

	// S3にアップロードします
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
// func PutFile(c context.Context, api S3ObjectAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
//     return api.PutObject(c, input)
// }

// func init() {

//     customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
//         if os.Getenv("AWS_LOCAL") == "true" {
//             return aws.Endpoint{
//                 PartitionID:   "aws",
//                 URL:           "http://localhost:4000",
//                 SigningRegion: "us-east-1",
//             }, nil
//         }
//         return aws.Endpoint{}, &aws.EndpointNotFoundError{}
//     })

//     cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolver(customResolver))
//     if err != nil {
//         panic(err)
//     }

//     client = s3.NewFromConfig(cfg, func(o *s3.Options) {
//         o.UsePathStyle = true
//     })
// }

// func Put(bucket string, filename string, file io.Reader) error {
//     input := &s3.PutObjectInput{
//         Bucket: aws.String(bucket),
//         Key:    aws.String(filename),
//         Body:   file,
//     }

//     _, err := PutFile(context.TODO(), client, input)
//     if err != nil {
//         return err
//     }

//     return nil
// }