package ceph

import (
	"gopkg.in/amz.v1/aws"
	"gopkg.in/amz.v1/s3"
	"myCloud/config"
)

var cephConn *s3.S3

func GetCephConnection() *s3.S3{
	if (cephConn!=nil){
		return cephConn
	}
	//init some information of ceph

	auth := aws.Auth{
		AccessKey: config.CephAccessKey,
		SecretKey: config.CephSecretKey,
	}

	curRegion:=aws.Region{
		Name:"default",
		EC2Endpoint: "http://47.102.123.183:7480",  //7480?
		S3Endpoint: "http://47.102.123.183:7480",
		S3BucketEndpoint: "",
		S3LocationConstraint: false,
		S3LowercaseBucket: false,
		Sign:aws.SignV2}

	//create an S3 style connection
	return s3.New(auth,curRegion)
}

func GetCephBucket(bucket string) *s3.Bucket{
	conn:=GetCephConnection()
	return conn.Bucket(bucket)
}