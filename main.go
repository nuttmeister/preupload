package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	bucket := flag.String("bucket", "", "The bucket to upload file to")
	key := flag.String("key", "", "The key to use for upload")
	expire := flag.Int("expire", 0, "Expiry of the presigned url in minutes")
	flag.Parse()
	if *bucket == "" || *key == "" || *expire == 0 {
		log.Fatal("-bucket, -key and -expire must be set")
	}

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	}

	svc := s3.New(cfg)

	req := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: bucket,
		Key:    key,
	})
	url, err := req.Presign(time.Duration(*expire) * time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("URL: %s", url)
}
